//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package worker

import (
	"context"
	"time"
)

// Options represents the configuration options for a Worker.
type Options struct {
	Domain       string
	BatchSize    int
	PollInterval time.Duration
	PollTimeout  time.Duration
	BaseContext  context.Context
}

func defaultOptions() Options {
	return Options{
		Domain:       "",
		BatchSize:    1,
		PollInterval: 100 * time.Millisecond,
		PollTimeout:  -1 * time.Millisecond,
	}
}

// Option defines a functional option for configuring a Worker.
type Option func(Options) Options

// WithBatchSize sets the number of tasks to fetch per poll for the worker.
func WithBatchSize(size int) Option {
	return func(o Options) Options {
		if size > 0 {
			o.BatchSize = size
		}
		return o
	}
}

// WithPollInterval sets the polling interval for the worker.
func WithPollInterval(interval time.Duration) Option {
	return func(o Options) Options {
		if interval > 0 {
			o.PollInterval = interval
		}
		return o
	}
}

// WithPollTimeout sets the polling timeout for the worker. Negative values mean server default.
func WithPollTimeout(timeout time.Duration) Option {
	return func(o Options) Options {
		o.PollTimeout = timeout
		return o
	}
}

// WithDomain sets the task domain for the worker.
func WithDomain(domain string) Option {
	return func(o Options) Options {
		o.Domain = domain
		return o
	}
}

// WithBaseContext sets the base context for the worker.
func WithBaseContext(ctx context.Context) Option {
	return func(o Options) Options {
		o.BaseContext = ctx
		return o
	}
}

func applyOptions(base Options, fns ...Option) Options {
	o := base
	for _, fn := range fns {
		if fn != nil {
			o = fn(o)
		}
	}
	return o
}
