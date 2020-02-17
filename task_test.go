package todotxt

import (
	"reflect"
	"testing"
	"time"
)

func TestTask_SetPriority(t *testing.T) {
	tests := []struct {
		Input byte
		IsErr bool
	}{
		{
			Input: 'A',
			IsErr: false,
		},
		{
			Input: 'A' - 1,
			IsErr: true,
		},
		{
			Input: 'Z',
			IsErr: false,
		},
		{
			Input: 'Z' + 1,
			IsErr: true,
		},
		{
			Input: 'a',
			IsErr: true,
		},
	}

	for _, test := range tests {
		task := &Task{}
		err := task.SetPriority(test.Input)

		if test.IsErr {
			if err == nil {
				t.Errorf("SetPriority(%s) should case error", string(test.Input))
			}
		} else {
			if test.Input != task.Priority() {
				t.Errorf("task Priority\ngot  %s\nwant %s", string(task.Priority()), string(test.Input))
			}
		}
	}
}

func TestTask_SetDescription(t *testing.T) {
	tests := []struct {
		Input       string
		Description string
		Projects    []string
		Contexts    []string
		Tags        map[string]string
	}{
		{
			Input:       "Hello World",
			Description: "Hello World",
		},
		{
			Input:       "Hello World +Project",
			Description: "Hello World +Project",
			Projects:    []string{"Project"},
		},
		{
			Input:       "Hello World @Context",
			Description: "Hello World @Context",
			Contexts:    []string{"Context"},
		},
		{
			Input:       "Hello World key:value",
			Description: "Hello World key:value",
			Tags:        map[string]string{"key": "value"},
		},
	}

	for _, test := range tests {
		task := &Task{}
		task.SetDescription(test.Input)

		if test.Description != task.Description() {
			t.Errorf("Description\ngot  %s\nwant %s", task.Description(), test.Description)
		}

		if !reflect.DeepEqual(test.Projects, task.Projects()) {
			t.Errorf("Projects\ngot  %s\nwant %s", task.Projects(), test.Projects)
		}

		if !reflect.DeepEqual(test.Contexts, task.Contexts()) {
			t.Errorf("Contexts\ngot  %s\nwant %s", task.Contexts(), test.Contexts)
		}

		if !reflect.DeepEqual(test.Tags, task.Tags()) {
			t.Errorf("Tags\ngot  %s\nwant %s", task.Tags(), test.Tags)
		}
	}
}

func TestTask_HasCreationDate(t *testing.T) {
	tests := []struct {
		Input  *Task
		Output bool
	}{
		{
			Input:  &Task{},
			Output: false,
		},
		{
			Input: &Task{
				CreationDate: time.Date(1, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			Output: true,
		},
	}

	for _, test := range tests {
		if test.Input.HasCreationDate() != test.Output {
			t.Errorf("task(%v).HasCreationDate got %v, want %v", test.Input, test.Input.HasCompletionDate(), test.Output)
		}
	}
}

func TestTask_HasCompletionDate(t *testing.T) {
	tests := []struct {
		Input  *Task
		Output bool
	}{
		{
			Input:  &Task{},
			Output: false,
		},
		{
			Input: &Task{
				CompletionDate: time.Date(1, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			Output: true,
		},
	}

	for _, test := range tests {
		if test.Input.HasCompletionDate() != test.Output {
			t.Errorf("task(%v).HasCompletionDate got %v, want %v", test.Input, test.Input.HasCompletionDate(), test.Output)
		}
	}
}
