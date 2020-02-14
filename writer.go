package todotxt

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"
)

// A Writer writes tasks using todo.txt encoding.
type Writer struct {
	w *bufio.Writer
}

// NewWriter returns New Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// Write writes single task to w.
// This method doesn't validate Task.
func (w *Writer) Write(t *Task) error {
	line := []string{}
	if t.Completed {
		line = append(line, "x")
	}

	if t.Priority != 0 {
		line = append(line, fmt.Sprintf("(%s)", string(t.Priority)))
	}

	if t.Completed && !reflect.DeepEqual(t.CompletionDate, time.Time{}) {
		line = append(line, t.CompletionDate.Format("2006-01-02"))
	}

	if !reflect.DeepEqual(t.CreationDate, time.Time{}) {
		line = append(line, t.CreationDate.Format("2006-01-02"))
	}

	if t.Description != "" {
		line = append(line, t.Description)
	}

	_, err := w.w.WriteString(strings.Join(line, " ") + "\n")
	return err
}

// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush() error {
	return w.w.Flush()
}

// WriteAll writes multiple tasks to w using Write and then calls Flush,
// returing any error from Write and Flush.
func (w *Writer) WriteAll(tasks []*Task) error {
	for _, t := range tasks {
		if err := w.Write(t); err != nil {
			return err
		}
	}
	return w.Flush()
}
