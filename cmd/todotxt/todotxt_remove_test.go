package main

import (
	"reflect"
	"testing"

	"github.com/kitagry/go-todotxt"
)

func TestRemoveTask(t *testing.T) {
	aTask := &todotxt.Task{}
	aTask.SetPriority('A')
	bTask := &todotxt.Task{}
	bTask.SetPriority('B')
	cTask := &todotxt.Task{}
	cTask.SetPriority('C')

	inputList := []*todotxt.Task{aTask, bTask, cTask}
	tests := []struct {
		Input  int
		Output []*todotxt.Task
	}{
		{
			Input:  -1,
			Output: []*todotxt.Task{aTask, bTask, cTask},
		},
		{
			Input:  0,
			Output: []*todotxt.Task{bTask, cTask},
		},
		{
			Input:  1,
			Output: []*todotxt.Task{aTask, cTask},
		},
		{
			Input:  2,
			Output: []*todotxt.Task{aTask, bTask},
		},
		{
			Input:  3,
			Output: []*todotxt.Task{aTask, bTask, cTask},
		},
	}

	for _, test := range tests {
		got := removeTask(inputList, test.Input)
		if !reflect.DeepEqual(got, test.Output) {
			t.Errorf("removeTask expected %v\ngot %v", test.Output, got)
		}
	}
}
