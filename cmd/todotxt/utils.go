package main

import (
	"errors"
	"os"

	"github.com/kitagry/go-todotxt"
	"golang.org/x/xerrors"
)

func getTasks(filename string) ([]*todotxt.Task, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("todo.txt is not found")
	}
	defer f.Close()

	r := todotxt.NewReader(f)
	tasks, err := r.ReadAll()
	if err != nil {
		return nil, xerrors.Errorf("Read todo error: %w", err)
	}
	return tasks, nil
}

func overWrite(filename string, tasks []*todotxt.Task) error {
	f, err := os.Create(filename)
	if err != nil {
		return errors.New("todo.txt is not found")
	}
	defer f.Close()

	w := todotxt.NewWriter(f)
	err = w.WriteAll(tasks)
	if err != nil {
		return xerrors.Errorf("Failed to write tasks: %w", err)
	}
	return nil
}
