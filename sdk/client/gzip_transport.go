// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package client

import (
	"compress/gzip"
	"io"
	"net/http"
)

// GzipTransport handles automatic gzip decompression
type GzipTransport struct {
	Transport http.RoundTripper
}

func (g *GzipTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add Accept-Encoding header to request gzip
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := g.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Check if response is gzipped
	if resp.Header.Get("Content-Encoding") == "gzip" {
		// Create gzip reader
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			if err = resp.Body.Close(); err != nil {
				return nil, err
			}
			return nil, err
		}

		// Replace body with decompressed reader
		resp.Body = &gzipReadCloser{
			Reader: reader,
			Closer: resp.Body,
		}

		// Remove Content-Encoding header so downstream doesn't try to decompress again
		resp.Header.Del("Content-Encoding")
		// Update Content-Length if present (it's now invalid after decompression)
		resp.Header.Del("Content-Length")
		resp.ContentLength = -1
	}

	return resp, nil
}

type gzipReadCloser struct {
	io.Reader
	io.Closer
}

func (g *gzipReadCloser) Read(p []byte) (n int, err error) {
	return g.Reader.Read(p)
}

func (g *gzipReadCloser) Close() error {
	if gzipReader, ok := g.Reader.(*gzip.Reader); ok {
		if err := gzipReader.Close(); err != nil {
			return err
		}
	}
	return g.Closer.Close()
}
