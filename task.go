package todotxt

import "time"

// Task represents a line struct for todo.txt format.
type Task struct {
	Completed bool

	// Priority is used for the next most important thing for you to get done
	// Priority is a uppercase character from A-Z.
	Priority byte

	CompletionDate time.Time

	CreationDate time.Time

	// Description is explanation of Task
	// Projects, Contexts and Tags are included.
	Description string

	Projects []string

	Contexts []string

	Tags map[string]string
}
