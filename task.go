package todotxt

import (
	"errors"
	"strings"
	"time"
)

// Task represents a line struct for todo.txt format.
type Task struct {
	Completed bool

	// priority is used for the next most important thing for you to get done
	// priority is a uppercase character from A-Z.
	priority byte

	CompletionDate time.Time

	CreationDate time.Time

	// description is explanation of Task
	// Projects, Contexts and Tags are included.
	//
	// It is because projects, contexts and tags can be placed anywhere in description.
	description string

	// projects is task's projects.
	// If you want to set projects, you can add tags to description, then use SetDescription.
	projects []string

	// contexts is task's contexts.
	// If you want to set contexts, you can add tags to description, then use SetDescription.
	contexts []string

	// tags is task's tags.
	// If you want to set tags, you can add tags to description, then use SetDescription.
	tags map[string]string
}

// SetPriority sets priority to task.
// this may returns error, when priority is not [A-Z].
func (t *Task) SetPriority(p byte) error {
	if p < 'A' || p > 'Z' {
		return errors.New("priority should be in [A-Z]")
	}
	t.priority = p
	return nil
}

// Priority returns task's priority.
func (t *Task) Priority() byte {
	return t.priority
}

// SetDescription set description, and parse descrition, then
// searches and sets projects, contexts and tags.
func (t *Task) SetDescription(d string) {
	t.description = d

	// check projects, contexts and tags
	ds := strings.Split(d, " ")
	for _, d := range ds {
		if len(d) == 1 {
			continue
		}

		if strings.HasPrefix(d, "+") {
			t.projects = append(t.projects, string(d[1:]))
		}

		if strings.HasPrefix(d, "@") {
			t.contexts = append(t.contexts, string(d[1:]))
		}

		if strings.Contains(d, ":") {
			i := strings.Index(d, ":")
			if i == 0 || i == len(d)-1 {
				continue
			}
			key := d[:i]
			value := d[i+1:]
			if t.tags == nil {
				t.tags = make(map[string]string)
			}
			t.tags[key] = value
		}
	}
}

// Description returns task's description.
func (t *Task) Description() string {
	return t.description
}

// Projects returns task's projects.
func (t *Task) Projects() []string {
	return t.projects
}

// Contexts returns task's contexts.
func (t *Task) Contexts() []string {
	return t.contexts
}

// Tags returns task's tags.
func (t *Task) Tags() map[string]string {
	return t.tags
}

// HasCreationDate returns bool whether task has CreationDate
func (t *Task) HasCreationDate() bool {
	return !t.CreationDate.IsZero()
}

// HasCompletionDate returns bool whether task has CompletionDate
func (t *Task) HasCompletionDate() bool {
	return !t.CompletionDate.IsZero()
}
