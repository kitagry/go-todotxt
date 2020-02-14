package todotxt

import (
	"reflect"
	"testing"
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
