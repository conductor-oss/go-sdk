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

func TestNewBaseWorker_Defaults(t *testing.T) {
	w := NewWorker("t", func(tk *model.Task) (interface{}, error) { return "ok", nil })
	assert.Equal(t, "t", w.TaskName())

	opts := w.Options()
	assert.Equal(t, 1, opts.BatchSize)
	assert.Equal(t, 100*time.Millisecond, opts.PollInterval)
	assert.Equal(t, -1*time.Millisecond, opts.PollTimeout)
	assert.Equal(t, "", opts.Domain)
}

func TestNewWorker_SetsHandlerAndName(t *testing.T) {
	called := false
	h := func(tk *model.Task) (interface{}, error) {
		called = true
		return "ok", nil
	}
	w := NewWorker("taskA", h)

	assert.Equal(t, "taskA", w.TaskName())
	handler := w.Handler()
	assert.NotNil(t, handler)
	out, err := handler(&model.Task{TaskDefName: "taskA"})
	assert.NoError(t, err)
	assert.Equal(t, "ok", out)
	assert.True(t, called)
}

func TestOptions_IgnoreInvalid_AndOrder(t *testing.T) {
	base := defaultOptions()

	// ignore non-positive
	opts := applyOptions(base, WithBatchSize(0))
	assert.Equal(t, 1, opts.BatchSize)
	opts = applyOptions(base, WithBatchSize(-5))
	assert.Equal(t, 1, opts.BatchSize)

	opts = applyOptions(base, WithPollInterval(0))
	assert.Equal(t, 100*time.Millisecond, opts.PollInterval)
	opts = applyOptions(base, WithPollInterval(-10*time.Millisecond))
	assert.Equal(t, 100*time.Millisecond, opts.PollInterval)

	// last wins
	opts = applyOptions(base, WithBatchSize(3), WithBatchSize(9))
	assert.Equal(t, 9, opts.BatchSize)

	opts = applyOptions(base, WithPollInterval(10*time.Millisecond), WithPollInterval(25*time.Millisecond))
	assert.Equal(t, 25*time.Millisecond, opts.PollInterval)

	opts = applyOptions(base, WithPollTimeout(0))
	assert.Equal(t, time.Duration(0), opts.PollTimeout)
	opts = applyOptions(base, WithPollTimeout(-1*time.Second))
	assert.Equal(t, -1*time.Second, opts.PollTimeout)
}

func TestNewWorker_NilHandler_AllowsNilButNotCalled(t *testing.T) {
	w := NewWorker("nilh", nil)
	assert.NotNil(t, w)
	handler := w.Handler()
	assert.Nil(t, handler)
}

func TestWorker_With_IsImmutable(t *testing.T) {
	w := NewWorker("t", func(tk *model.Task) (interface{}, error) { return nil, nil })
	w2 := w.With(WithBatchSize(2))
	_ = w2.With(WithDomain("d2"))

	opts := w.Options()
	assert.Equal(t, 1, opts.BatchSize)
	assert.Equal(t, "", opts.Domain)
}
