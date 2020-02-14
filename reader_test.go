package todotxt_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/kitagry/go-todotxt"
)

func TestReader_Read(t *testing.T) {
	tests := []struct {
		Input  string
		Output *todotxt.Task
	}{
		{
			Input: "x (A) 2020-01-15 2020-01-01 measure space for +chapelShelving @chapel due:2020-01-12",
			Output: &todotxt.Task{
				Completed:      true,
				Priority:       'A',
				CompletionDate: time.Date(2020, 01, 15, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
				Description:    "measure space for +chapelShelving @chapel due:2020-01-12",
				Projects:       []string{"chapelShelving"},
				Contexts:       []string{"chapel"},
				Tags: map[string]string{
					"due": "2020-01-12",
				},
			},
		},
		{
			Input: "(A) 2020-01-01",
			Output: &todotxt.Task{
				Priority:     'A',
				Completed:    false,
				CreationDate: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			Input: "(a) 2020-01-01",
			Output: &todotxt.Task{
				Description: "(a) 2020-01-01",
			},
		},
		{
			Input: "x (B) 2020-01-01",
			Output: &todotxt.Task{
				Priority:       'B',
				Completed:      true,
				CompletionDate: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			Input: "xylophone lesson",
			Output: &todotxt.Task{
				Completed:   false,
				Description: "xylophone lesson",
			},
		},
		{
			Input: "x 2020-01-15 2020-01-01",
			Output: &todotxt.Task{
				Completed:      true,
				CompletionDate: time.Date(2020, 01, 15, 0, 0, 0, 0, time.UTC),
				CreationDate:   time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			Input: "(A) Thank Mom for the meatballs @phone",
			Output: &todotxt.Task{
				Priority:    'A',
				Description: "Thank Mom for the meatballs @phone",
				Contexts:    []string{"phone"},
			},
		},
		{
			Input: "Post signs around the neighborhood +GarageSale",
			Output: &todotxt.Task{
				Description: "Post signs around the neighborhood +GarageSale",
				Projects:    []string{"GarageSale"},
			},
		},
		{
			Input: "@GroceryStore Eskimo pies",
			Output: &todotxt.Task{
				Description: "@GroceryStore Eskimo pies",
				Contexts:    []string{"GroceryStore"},
			},
		},
		{
			Input: "Email SoAndSo at soandso@example.com",
			Output: &todotxt.Task{
				Description: "Email SoAndSo at soandso@example.com",
			},
		},
		{
			Input: "Learn how to add 2+2",
			Output: &todotxt.Task{
				Description: "Learn how to add 2+2",
			},
		},
		{
			Input: "X 2012-01-01 Make resolutions",
			Output: &todotxt.Task{
				Description: "X 2012-01-01 Make resolutions",
			},
		},
		{
			Input: "1 + 1",
			Output: &todotxt.Task{
				Description: "1 + 1",
			},
		},
		{
			Input: "hello @ world",
			Output: &todotxt.Task{
				Description: "hello @ world",
			},
		},
		{
			Input: "a: b",
			Output: &todotxt.Task{
				Description: "a: b",
			},
		},
		{
			Input: "a :b",
			Output: &todotxt.Task{
				Description: "a :b",
			},
		},
	}

	for _, test := range tests {
		r := todotxt.NewReader(strings.NewReader(test.Input))
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

	output := []*todotxt.Task{
		{
			Priority:    'A',
			Description: "Thank Mom for the meatballs @phone",
			Contexts:    []string{"phone"},
		},
		{
			Priority:    'B',
			Description: "Schedule Goodwill pickup +GarageSale @phone",
			Projects:    []string{"GarageSale"},
			Contexts:    []string{"phone"},
		},
		{
			Description: "Post signs around the neighborhood +GarageSale",
			Projects:    []string{"GarageSale"},
		},
		{
			Description: "@GroceryStore Eskimo pies",
			Contexts:    []string{"GroceryStore"},
		},
	}

	r := todotxt.NewReader(strings.NewReader(input))
	tasks, err := r.ReadAll()
	if err != nil {
		t.Errorf("ReadAll should be succeed: %v", err)
	}

	if !reflect.DeepEqual(tasks, output) {
		t.Errorf("ReadAll error\ngot %v\nwant %v", tasks, output)
	}
}
