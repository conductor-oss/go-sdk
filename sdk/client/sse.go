//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"strings"
)

// SSEEvent is one server-sent event.
type SSEEvent struct {
	// Event is the event name, from an "event:" line. Empty for unnamed events.
	Event string
	// Data is the payload, with multiple "data:" lines joined by newlines.
	Data string
	// ID is the last-event id, from an "id:" line.
	ID string
}

// StreamSSE opens a server-sent event stream and delivers events on a channel.
//
// This does not go through APIClient.Get. That path sets Accept:
// application/json and reads the entire response body before returning, so on
// a stream it would block until the run ended and then deliver everything at
// once. The request is built here instead, but shares the APIClient's token so
// authentication is minted and cached once for the process.
//
// The channel closes when the stream ends, the context is cancelled, or an
// error occurs; call Err after the channel closes to see which.
func (c *APIClient) StreamSSE(ctx context.Context, path, lastEventID string) (<-chan SSEEvent, *SSEStream, error) {
	req, err := c.newStreamRequest(ctx, path, lastEventID)
	if err != nil {
		return nil, nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("open sse stream: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err := fmt.Errorf("open sse stream: unexpected status %d", resp.StatusCode)
		if cerr := resp.Body.Close(); cerr != nil {
			err = fmt.Errorf("%w (and closing the body: %v)", err, cerr)
		}
		return nil, nil, err
	}

	stream := &SSEStream{resp: resp}
	ch := make(chan SSEEvent)
	go stream.read(ctx, ch)
	return ch, stream, nil
}

// SSEStream owns the response body of a live stream.
type SSEStream struct {
	resp *http.Response
	err  error
}

// Err reports why the stream ended, or nil for a clean end.
func (s *SSEStream) Err() error { return s.err }

// Close ends the stream early.
func (s *SSEStream) Close() error { return s.resp.Body.Close() }

// read parses the wire format: lines of "field: value", with a blank line
// terminating each event. Lines beginning with ':' are comments, which servers
// use as keep-alives.
func (s *SSEStream) read(ctx context.Context, ch chan<- SSEEvent) {
	defer close(ch)
	// A close failure after the stream ended is the only error left to report,
	// so it becomes Err() when nothing else already has.
	defer func() {
		if err := s.resp.Body.Close(); err != nil && s.err == nil {
			s.err = err
		}
	}()

	scanner := bufio.NewScanner(s.resp.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)

	var ev SSEEvent
	var data []string

	flush := func() bool {
		if ev.Event == "" && len(data) == 0 {
			return true
		}
		ev.Data = strings.Join(data, "\n")
		select {
		case ch <- ev:
		case <-ctx.Done():
			s.err = ctx.Err()
			return false
		}
		ev, data = SSEEvent{}, nil
		return true
	}

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "":
			if !flush() {
				return
			}
		case strings.HasPrefix(line, ":"):
			// keep-alive comment
		default:
			applyField(line, &ev, &data)
		}
	}
	flush()
	if err := scanner.Err(); err != nil && s.err == nil {
		s.err = err
	}
}

// applyField parses one "field: value" line into the event being assembled.
// Unknown fields are ignored, as the SSE spec requires.
func applyField(line string, ev *SSEEvent, data *[]string) {
	field, value, found := strings.Cut(line, ":")
	if !found {
		field, value = line, ""
	}
	value = strings.TrimPrefix(value, " ")
	switch field {
	case "event":
		ev.Event = value
	case "data":
		*data = append(*data, value)
	case "id":
		ev.ID = value
	}
}

// newStreamRequest builds the streaming request, stamping the same token the
// APIClient uses so the stream and the JSON calls share one authentication.
func (c *APIClient) newStreamRequest(ctx context.Context, path, lastEventID string) (*http.Request, error) {
	url := strings.TrimSuffix(c.httpRequester.httpSettings.BaseUrl, "/") + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	if lastEventID != "" {
		req.Header.Set("Last-Event-ID", lastEventID)
	}
	// Reuse the APIClient's token manager so the stream shares the cached
	// token rather than minting a second one.
	if tm := c.httpRequester.tokenManager; tm != nil {
		if token, err := tm.RefreshToken(
			c.httpRequester.httpSettings, c.httpRequester.httpClient); err == nil && token != "" {
			req.Header.Set("X-Authorization", token)
		}
	}
	return req, nil
}
