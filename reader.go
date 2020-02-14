package todotxt

import (
	"bufio"
	"io"
	"regexp"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

var (
	priorityRegex = regexp.MustCompile(`^\([A-Z]\)`)
	dateRegex     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)
)

// Reader reads records from a todo.txt style text.
type Reader struct {
	r *bufio.Reader
}

// NewReader returns a Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

func (r *Reader) parseTask(line string) (*Task, error) {
	task := &Task{}

	// check completed
	if strings.HasPrefix(line, "x ") {
		task.Completed = true
		line = strings.Trim(string(line[1:]), " ")
	}

	// check priority
	if priorityRegex.MatchString(line) {
		err := task.SetPriority(line[1])
		if err != nil {
			return nil, xerrors.Errorf("SetPriority failed: %w", err)
		}
		line = strings.Trim(string(line[3:]), " ")
	}

	// check completion date
	if task.Completed && dateRegex.MatchString(line) {
		c, err := time.Parse("2006-01-02", string(line[:10]))
		if err != nil {
			return nil, xerrors.Errorf("CompletionDate parse error: %w", err)
		}
		task.CompletionDate = c
		line = strings.Trim(string(line[10:]), " ")
	}

	// check creation date
	if dateRegex.MatchString(line) {
		c, err := time.Parse("2006-01-02", string(line[:10]))
		if err != nil {
			return nil, xerrors.Errorf("CreationDate parse error: %w", err)
		}
		task.CreationDate = c
		line = strings.Trim(string(line[10:]), " ")
	}

	task.SetDescription(line)

	return task, nil
}

// Read reads one Task from r.
func (r *Reader) Read() (*Task, error) {
	b, _, err := r.r.ReadLine()
	if err != nil {
		return nil, xerrors.Errorf("ReadLine return error: %w", err)
	}

	task, err := r.parseTask(string(b))
	if err != nil {
		return nil, xerrors.Errorf("parseTask return error: %w", err)
	}
	return task, nil
}

// ReadAll reads all Tasks from r.
func (r *Reader) ReadAll() (tasks []*Task, err error) {
	for {
		task, err := r.Read()
		if xerrors.Is(err, io.EOF) {
			return tasks, nil
		}

		if err != nil {
			return nil, xerrors.Errorf("Read returns error: %w", err)
		}
		tasks = append(tasks, task)
	}
}
