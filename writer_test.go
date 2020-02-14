package todotxt

import (
	"bytes"
	"testing"
	"time"
)

func TestWriter_Write(t *testing.T) {
	tests := []struct {
		Input  *Task
		Output string
	}{
		{
			Input:  &Task{},
			Output: "\n",
		},
		{
			Input: &Task{
				priority:  'A',
				Completed: true,
			},
			Output: "x (A)\n",
		},
		{
			Input: &Task{
				priority:       'A',
				Completed:      true,
				CompletionDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Output: "x (A) 2020-01-01 2020-01-01\n",
		},
		{
			Input: &Task{
				priority:       'A',
				Completed:      false,
				CompletionDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Output: "(A) 2020-01-01\n",
		},
		{
			Input: &Task{
				priority:       'A',
				Completed:      true,
				CompletionDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				description:    "Hello World",
			},
			Output: "x (A) 2020-01-01 2020-01-01 Hello World\n",
		},
		{
			Input: &Task{
				priority:       'A',
				Completed:      true,
				CompletionDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				description:    "Hello World +Project @context key:value",
			},
			Output: "x (A) 2020-01-01 2020-01-01 Hello World +Project @context key:value\n",
		},
	}

	for _, test := range tests {
		buf := &bytes.Buffer{}
		w := NewWriter(buf)
		err := w.Write(test.Input)
		if err != nil {
			t.Errorf("Writer Write should success: %v", err)
			break
		}

		err = w.Flush()
		if err != nil {
			t.Errorf("Writer Flush should success: %v", err)
			break
		}

		result := buf.String()
		if result != test.Output {
			t.Errorf("Writer result\ngot  %s\nwant %s", result, test.Output)
		}
	}
}

func TestWriter_WriteAll(t *testing.T) {
	input := []*Task{
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

	output := `(A) Thank Mom for the meatballs @phone
(B) Schedule Goodwill pickup +GarageSale @phone
Post signs around the neighborhood +GarageSale
@GroceryStore Eskimo pies
`
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	err := w.WriteAll(input)
	if err != nil {
		t.Errorf("WriteAll return error: %v", err)
		return
	}

	result := buf.String()
	if result != output {
		t.Errorf("Writer WriteAll\n got  %s\nwant %s", result, output)
	}
}
