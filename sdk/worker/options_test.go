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
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestOption_WithBatchSize_SetsOnlyPositive(t *testing.T) {
	base := Options{BatchSize: 99}

	opts := WithBatchSize(-1)(base)
	assert.Equal(t, 99, opts.BatchSize, "negatives ignored")

	opts = WithBatchSize(0)(base)
	assert.Equal(t, 99, opts.BatchSize, "zero ignored")

	opts = WithBatchSize(1)(base)
	assert.Equal(t, 1, opts.BatchSize)

	opts = WithBatchSize(5)(base)
	assert.Equal(t, 5, opts.BatchSize)
}

func TestOption_WithPollInterval_SetsOnlyPositive(t *testing.T) {
	base := Options{PollInterval: 987 * time.Millisecond}

	opts := WithPollInterval(-10 * time.Millisecond)(base)
	assert.Equal(t, 987*time.Millisecond, opts.PollInterval, "negatives ignored")

	opts = WithPollInterval(0)(base)
	assert.Equal(t, 987*time.Millisecond, opts.PollInterval, "zero ignored")

	opts = WithPollInterval(50 * time.Millisecond)(base)
	assert.Equal(t, 50*time.Millisecond, opts.PollInterval)
}

func TestOption_WithPollTimeout_SetsAnyValue(t *testing.T) {
	base := Options{PollTimeout: 111 * time.Millisecond}

	opts := WithPollTimeout(0)(base)
	assert.Equal(t, 0*time.Millisecond, opts.PollTimeout)

	opts = WithPollTimeout(-1 * time.Second)(base)
	assert.Equal(t, -1*time.Second, opts.PollTimeout)

	opts = WithPollTimeout(250 * time.Millisecond)(base)
	assert.Equal(t, 250*time.Millisecond, opts.PollTimeout)
}

func TestOption_WithDomain_SetsAsIs(t *testing.T) {
	base := Options{Domain: "old"}

	opts := WithDomain("")(base)
	assert.Equal(t, "", opts.Domain)

	opts = WithDomain("testing")(base)
	assert.Equal(t, "testing", opts.Domain)

	opts = WithDomain("  spaced  ")(base)
	assert.Equal(t, "  spaced  ", opts.Domain, "no trimming in option")
}

func TestOptions_Composition_OrderMatters_AndBaseNotMutated(t *testing.T) {
	base := defaultOptions()

	opts := applyOptions(base, WithBatchSize(1), WithBatchSize(10))
	assert.Equal(t, 10, opts.BatchSize)

	opts = applyOptions(base, WithPollInterval(10*time.Millisecond), WithPollInterval(25*time.Millisecond))
	assert.Equal(t, 25*time.Millisecond, opts.PollInterval)

	assert.Equal(t, defaultOptions(), base)
}

func TestApplyOptions_WithNilOptions(t *testing.T) {
	base := defaultOptions()
	opts := applyOptions(base, nil, WithBatchSize(5), nil, WithDomain("test"))

	assert.Equal(t, 5, opts.BatchSize)
	assert.Equal(t, "test", opts.Domain)

	assert.Equal(t, defaultOptions(), base)
}

func TestDefaultOptions(t *testing.T) {
	opts := defaultOptions()
	assert.Equal(t, "", opts.Domain)
	assert.Equal(t, 1, opts.BatchSize)
	assert.Equal(t, 100*time.Millisecond, opts.PollInterval)
	assert.Equal(t, -1*time.Millisecond, opts.PollTimeout)
	assert.Nil(t, opts.BaseContext)
}

func TestNewWorker_AppliesOptions(t *testing.T) {
	ctx := context.WithValue(context.Background(), "base", true)

	w := NewWorker(
		"opt_task",
		func(tk *model.Task) (any, error) { return "ok", nil },
		WithBatchSize(3),
		WithPollInterval(123*time.Millisecond),
		WithPollTimeout(-1*time.Second),
		WithDomain("testing"),
		WithBaseContext(ctx),
	)

	opts := w.Options()
	assert.Equal(t, 3, opts.BatchSize)
	assert.Equal(t, 123*time.Millisecond, opts.PollInterval)
	assert.Equal(t, -1*time.Second, opts.PollTimeout)
	assert.Equal(t, "testing", opts.Domain)
	assert.Equal(t, ctx, opts.BaseContext)
}

func TestNewTypedWorker_AppliesOptions(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct {
		B int `json:"b"`
	}

	ctx := context.WithValue(context.Background(), "typed", true)

	tw := NewTypedWorker[In, Out](
		"typed_opt",
		func(ctx TaskContext, in In) (Out, error) { return Out{B: in.A}, nil },
		WithBatchSize(7),
		WithPollInterval(42*time.Millisecond),
		WithPollTimeout(5*time.Second),
		WithDomain("typed-domain"),
		WithBaseContext(ctx),
	)

	opts := tw.Options()
	assert.Equal(t, 7, opts.BatchSize)
	assert.Equal(t, 42*time.Millisecond, opts.PollInterval)
	assert.Equal(t, 5*time.Second, opts.PollTimeout)
	assert.Equal(t, "typed-domain", opts.Domain)
	assert.Equal(t, ctx, opts.BaseContext)

	w := tw.With()
	assert.NotSame(t, tw, w)
	wOpts := w.Options()
	assert.Equal(t, 7, wOpts.BatchSize)
	assert.Equal(t, 42*time.Millisecond, wOpts.PollInterval)
	assert.Equal(t, 5*time.Second, wOpts.PollTimeout)
	assert.Equal(t, "typed-domain", wOpts.Domain)
	assert.Equal(t, ctx, wOpts.BaseContext)
}

func TestApplyOptions_Immutability(t *testing.T) {
	base := defaultOptions()
	before := base

	_ = applyOptions(base,
		WithBatchSize(99),
		WithDomain("x"),
		WithPollInterval(5*time.Second),
		WithPollTimeout(0),
	)

	assert.Equal(t, before, base, "base options must remain unchanged")
}
