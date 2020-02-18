package todotxt

import (
	"bufio"
	"io"
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
	_, err := w.w.WriteString(t.Format() + "\n")
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
