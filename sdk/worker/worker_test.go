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
	w := newBaseWorker("t")
	assert.Equal(t, "t", w.taskName)
	assert.Equal(t, 1, w.batchSize)
	assert.Equal(t, 500*time.Millisecond, w.pollInterval)
	assert.Equal(t, -1*time.Millisecond, w.pollTimeout)
	assert.Equal(t, "", w.domain)

	_, isJSON := w.binder.(JSONBinder)
	assert.True(t, isJSON)
}

func TestNewWorker_SetsHandlerAndName(t *testing.T) {
	called := false
	h := func(tk *model.Task) (interface{}, error) {
		called = true
		return "ok", nil
	}
	w := NewWorker("taskA", h)

	assert.Equal(t, "taskA", w.taskName)
	out, err := w.handler(&model.Task{TaskDefName: "taskA"})
	assert.NoError(t, err)
	assert.Equal(t, "ok", out)
	assert.True(t, called)
}

func TestOptions_IgnoreInvalid_AndOrder(t *testing.T) {
	w := newBaseWorker("t")

	// ignore non-positive
	WithBatchSize(0)(w)
	assert.Equal(t, 1, w.batchSize)
	WithBatchSize(-5)(w)
	assert.Equal(t, 1, w.batchSize)

	WithPollInterval(0)(w)
	assert.Equal(t, 500*time.Millisecond, w.pollInterval)
	WithPollInterval(-10 * time.Millisecond)(w)
	assert.Equal(t, 500*time.Millisecond, w.pollInterval)

	// last wins
	WithBatchSize(3)(w)
	WithBatchSize(9)(w)
	assert.Equal(t, 9, w.batchSize)

	WithPollInterval(10 * time.Millisecond)(w)
	WithPollInterval(25 * time.Millisecond)(w)
	assert.Equal(t, 25*time.Millisecond, w.pollInterval)

	WithPollTimeout(0)(w)
	assert.Equal(t, time.Duration(0), w.pollTimeout)
	WithPollTimeout(-1 * time.Second)(w)
	assert.Equal(t, -1*time.Second, w.pollTimeout)
}

func TestWorker_ProviderInterface_ReturnsSelf(t *testing.T) {
	w := NewWorker("p", func(*model.Task) (interface{}, error) { return nil, nil })
	var p Provider = w
	got := p.Worker()
	assert.Same(t, w, got)
}

func TestNewWorker_NilHandler_AllowsNilButNotCalled(t *testing.T) {
	w := NewWorker("nilh", nil)
	assert.NotNil(t, w)
	assert.Nil(t, w.handler)
}
