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
	"errors"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/stretchr/testify/assert"
)

func TestTypedWorker_JSONBinding_Success_DefaultBinder(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct {
		B int `json:"b"`
	}

	tw := NewTypedWorker(
		"typed",
		func(ctx TaskContext, in In) (Out, error) {
			assert.NotNil(t, ctx)
			return Out{B: in.A + 10}, nil
		},
	)

	res, err := tw.adapter()(&model.Task{
		TaskDefName: "typed",
		InputData:   map[string]any{"a": 5},
	})
	assert.NoError(t, err)
	assert.Equal(t, Out{B: 15}, res)
}

func TestTypedWorker_JSONBinding_Error_DefaultBinder(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct{}

	tw := NewTypedWorker(
		"typed_err",
		func(ctx TaskContext, in In) (Out, error) { return Out{}, nil },
	)

	_, err := tw.adapter()(&model.Task{
		TaskDefName: "typed_err",
		InputData:   map[string]any{"a": "not_an_int"},
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "input binding error")
	assert.Contains(t, err.Error(), "typed_err")
}

func TestTypedWorker_HandlerError_Propagates(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	want := errors.New("handler failed")

	tw := NewTypedWorker("t", func(ctx TaskContext, _ In) (struct{}, error) {
		return struct{}{}, want
	})

	_, err := tw.adapter()(&model.Task{
		TaskDefName: "t",
		InputData:   map[string]any{"a": 1},
	})

	assert.ErrorIs(t, err, want)
}

func TestTypedWorker_Constructors_And_With(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct {
		B int `json:"b"`
	}

	tw := NewTypedWorker("t", func(ctx TaskContext, in In) (Out, error) {
		return Out{B: in.A}, nil
	})

	w := tw.With()
	assert.NotNil(t, w)
	assert.NotSame(t, tw, w)

	handler := w.Handler()
	assert.NotNil(t, handler)
	res, err := handler(&model.Task{
		TaskDefName: "t",
		InputData:   map[string]any{"a": 7},
	})
	assert.NoError(t, err)
	assert.Equal(t, Out{B: 7}, res)
}

func TestTypedWorker_WithTaskContext_Variant(t *testing.T) {
	type In struct{}
	type Out struct{}

	seen := false
	tw := NewTypedWorker("t", func(ctx TaskContext, _ In) (Out, error) {
		seen = true
		return Out{}, nil
	})

	_, err := tw.adapter()(&model.Task{TaskDefName: "t"})
	assert.NoError(t, err)
	assert.True(t, seen)
}

func TestAdapter_TypedNilPointerInOut(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct{ X int }

	tw := NewTypedWorker("t", func(ctx TaskContext, in In) (*Out, error) {
		return nil, nil
	})

	v, err := tw.adapter()(&model.Task{TaskDefName: "t", InputData: map[string]any{"a": 0}})
	assert.NoError(t, err)
	assert.Nil(t, v)
}
func TestAdapter_NonNilPointerInOut(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct{ X int }

	tw := NewTypedWorker("t", func(ctx TaskContext, in In) (*Out, error) {
		return &Out{X: 42}, nil
	})

	v, err := tw.adapter()(&model.Task{TaskDefName: "t", InputData: map[string]any{"a": 1}})
	assert.NoError(t, err)
	out, ok := v.(*Out)
	assert.True(t, ok)
	assert.Equal(t, 42, out.X)
}
func TestAdapter_NilInputData_IsSafe(t *testing.T) {
	type In struct {
		A int `json:"a"`
	}
	type Out struct {
		B int `json:"b"`
	}

	tw := NewTypedWorker("t", func(ctx TaskContext, in In) (Out, error) {
		return Out{B: in.A + 1}, nil
	})
	res, err := tw.adapter()(&model.Task{TaskDefName: "t", InputData: nil})
	assert.NoError(t, err)
	assert.Equal(t, Out{B: 1}, res)
}

func TestTypedWorker_Options(t *testing.T) {
	type In struct{}
	type Out struct{}

	tw := NewTypedWorker("test", func(ctx TaskContext, in In) (Out, error) {
		return Out{}, nil
	})

	opts := tw.Options()
	assert.Equal(t, "test", tw.TaskName())
	assert.Equal(t, 1, opts.BatchSize)
	assert.Equal(t, 100*time.Millisecond, opts.PollInterval)
	assert.Equal(t, -1*time.Millisecond, opts.PollTimeout)
	assert.Equal(t, "", opts.Domain)
}

func TestTypedWorker_WithOptions(t *testing.T) {
	type In struct{}
	type Out struct{}

	tw := NewTypedWorker("test", func(ctx TaskContext, in In) (Out, error) {
		return Out{}, nil
	})

	tw2 := tw.With(WithBatchSize(5), WithDomain("test-domain"))

	opts1 := tw.Options()
	assert.Equal(t, 1, opts1.BatchSize)
	assert.Equal(t, "", opts1.Domain)

	opts2 := tw2.Options()
	assert.Equal(t, 5, opts2.BatchSize)
	assert.Equal(t, "test-domain", opts2.Domain)

	tw2Concrete := tw2.(*TypedWorker[In, Out])
	assert.NotSame(t, tw.handler, tw2Concrete.handler)
}
