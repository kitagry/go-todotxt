package todotxt

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestReader_Read(t *testing.T) {
	tests := []struct {
		Input  string
		Output *Task
	}{
		{
			Input: "x (A) 2020-01-15 2020-01-01 measure space for +chapelShelving @chapel due:2020-01-12",
			Output: &Task{
				Completed:      true,
				priority:       'A',
				CompletionDate: time.Date(2020, 01, 15, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
				description:    "measure space for +chapelShelving @chapel due:2020-01-12",
				projects:       []string{"chapelShelving"},
				contexts:       []string{"chapel"},
				tags: map[string]string{
					"due": "2020-01-12",
				},
			},
		},
		{
			Input: "(A) 2020-01-01",
			Output: &Task{
				priority:     'A',
				Completed:    false,
				CreationDate: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			Input: "(a) 2020-01-01",
			Output: &Task{
				description: "(a) 2020-01-01",
			},
		},
		{
			Input: "x (B) 2020-01-01",
			Output: &Task{
				priority:       'B',
				Completed:      true,
				CompletionDate: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			Input: "xylophone lesson",
			Output: &Task{
				Completed:   false,
				description: "xylophone lesson",
			},
		},
		{
			Input: "x 2020-01-15 2020-01-01",
			Output: &Task{
				Completed:      true,
				CompletionDate: time.Date(2020, 01, 15, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			Input: "(A) Thank Mom for the meatballs @phone",
			Output: &Task{
				priority:    'A',
				description: "Thank Mom for the meatballs @phone",
				contexts:    []string{"phone"},
			},
		},
		{
			Input: "Post signs around the neighborhood +GarageSale",
			Output: &Task{
				description: "Post signs around the neighborhood +GarageSale",
				projects:    []string{"GarageSale"},
			},
		},
		{
			Input: "@GroceryStore Eskimo pies",
			Output: &Task{
				description: "@GroceryStore Eskimo pies",
				contexts:    []string{"GroceryStore"},
			},
		},
		{
			Input: "Email SoAndSo at soandso@example.com",
			Output: &Task{
				description: "Email SoAndSo at soandso@example.com",
			},
		},
		{
			Input: "Learn how to add 2+2",
			Output: &Task{
				description: "Learn how to add 2+2",
			},
		},
		{
			Input: "X 2012-01-01 Make resolutions",
			Output: &Task{
				description: "X 2012-01-01 Make resolutions",
			},
		},
		{
			Input: "1 + 1",
			Output: &Task{
				description: "1 + 1",
			},
		},
		{
			Input: "hello @ world",
			Output: &Task{
				description: "hello @ world",
			},
		},
		{
			Input: "a: b",
			Output: &Task{
				description: "a: b",
			},
		},
		{
			Input: "a :b",
			Output: &Task{
				description: "a :b",
			},
		},
	}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.Input))
		task, err := r.Read()
		if err != nil {
			t.Errorf("Read should success: %v", err)
			return
		}

		if !reflect.DeepEqual(task, test.Output) {
			t.Errorf("Read() output:\ngot %v\n want %v", task, test.Output)
			return
		}
	}
}

func TestReader_ReadAll(t *testing.T) {
	input := `(A) Thank Mom for the meatballs @phone
(B) Schedule Goodwill pickup +GarageSale @phone
Post signs around the neighborhood +GarageSale
@GroceryStore Eskimo pies
`

	output := []*Task{
		{
			priority:    'A',
			description: "Thank Mom for the meatballs @phone",
			contexts:    []string{"phone"},
		},
		{
			priority:    'B',
			description: "Schedule Goodwill pickup +GarageSale @phone",
			projects:    []string{"GarageSale"},
			contexts:    []string{"phone"},
		},
		{
			description: "Post signs around the neighborhood +GarageSale",
			projects:    []string{"GarageSale"},
		},
		{
			description: "@GroceryStore Eskimo pies",
			contexts:    []string{"GroceryStore"},
		},
	}

	r := NewReader(strings.NewReader(input))
	tasks, err := r.ReadAll()
	if err != nil {
		t.Errorf("ReadAll should be succeed: %v", err)
	}

	if !reflect.DeepEqual(tasks, output) {
		t.Errorf("ReadAll error\ngot %v\nwant %v", tasks, output)
	}
}
