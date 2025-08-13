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
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestOption_WithBatchSize_SetsOnlyPositive(t *testing.T) {
	var w Worker
	w.batchSize = 99

	WithBatchSize(-1)(&w)
	assert.Equal(t, 99, w.batchSize, "negatives ignored")

	WithBatchSize(0)(&w)
	assert.Equal(t, 99, w.batchSize, "zero ignored")

	WithBatchSize(1)(&w)
	assert.Equal(t, 1, w.batchSize)

	WithBatchSize(5)(&w)
	assert.Equal(t, 5, w.batchSize, "last wins")
}

func TestOption_WithPollInterval_SetsOnlyPositive(t *testing.T) {
	var w Worker
	w.pollInterval = 987 * time.Millisecond

	WithPollInterval(-10 * time.Millisecond)(&w)
	assert.Equal(t, 987*time.Millisecond, w.pollInterval, "negatives ignored")

	WithPollInterval(0)(&w)
	assert.Equal(t, 987*time.Millisecond, w.pollInterval, "zero ignored")

	WithPollInterval(50 * time.Millisecond)(&w)
	assert.Equal(t, 50*time.Millisecond, w.pollInterval)
}

func TestOption_WithPollTimeout_SetsAnyValue(t *testing.T) {
	var w Worker
	w.pollTimeout = 111 * time.Millisecond

	WithPollTimeout(0)(&w) // допускается
	assert.Equal(t, 0*time.Millisecond, w.pollTimeout)

	WithPollTimeout(-1 * time.Second)(&w) // «server default» по договорённости
	assert.Equal(t, -1*time.Second, w.pollTimeout)

	WithPollTimeout(250 * time.Millisecond)(&w)
	assert.Equal(t, 250*time.Millisecond, w.pollTimeout)
}

func TestOption_WithDomain_SetsAsIs(t *testing.T) {
	var w Worker
	w.domain = "old"

	WithDomain("")(&w)
	assert.Equal(t, "", w.domain)

	WithDomain("testing")(&w)
	assert.Equal(t, "testing", w.domain)

	WithDomain("  spaced  ")(&w)
	assert.Equal(t, "  spaced  ", w.domain, "no trimming in option")
}

func TestOptions_Composition_OrderMatters(t *testing.T) {
	var w Worker
	WithBatchSize(1)(&w)
	WithBatchSize(10)(&w) // последняя выигрывает
	assert.Equal(t, 10, w.batchSize)

	WithPollInterval(10 * time.Millisecond)(&w)
	WithPollInterval(25 * time.Millisecond)(&w)
	assert.Equal(t, 25*time.Millisecond, w.pollInterval)
}

func TestNewWorker_AppliesOptions(t *testing.T) {
	w := NewWorker(
		"opt_task",
		func(tk *model.Task) (any, error) { return "ok", nil },
		WithBatchSize(3),
		WithPollInterval(123*time.Millisecond),
		WithPollTimeout(-1*time.Second),
		WithDomain("testing"),
	)

	assert.Equal(t, 3, w.batchSize)
	assert.Equal(t, 123*time.Millisecond, w.pollInterval)
	assert.Equal(t, -1*time.Second, w.pollTimeout)
	assert.Equal(t, "testing", w.domain)
}

func TestNewTypedWorker_AppliesOptions_ToBase(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct {
		B int `json:"b"`
	}

	tw := NewTypedWorker[In, Out](
		"typed_opt",
		func(ctx TaskContext, in In) (Out, error) { return Out{B: in.A}, nil },
		WithBatchSize(7),
		WithPollInterval(42*time.Millisecond),
		WithPollTimeout(5*time.Second),
		WithDomain("typed-domain"),
	)

	assert.Equal(t, 7, tw.base.batchSize)
	assert.Equal(t, 42*time.Millisecond, tw.base.pollInterval)
	assert.Equal(t, 5*time.Second, tw.base.pollTimeout)
	assert.Equal(t, "typed-domain", tw.base.domain)

	w := tw.Worker()
	assert.NotSame(t, tw.base, w)
	assert.Equal(t, 7, w.batchSize)
	assert.Equal(t, 42*time.Millisecond, w.pollInterval)
	assert.Equal(t, 5*time.Second, w.pollTimeout)
	assert.Equal(t, "typed-domain", w.domain)
}
