package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/kitagry/go-todotxt"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var undoFlags = []cli.Flag{
	&cli.StringFlag{Name: "file", Value: "todo.txt", Aliases: []string{"f"}, Usage: "Path to todo.txt file"},
}

func todotxtUndo(c *cli.Context) error {
	todotxtFile := c.String("file")
	f, err := os.Open(todotxtFile)
	if err != nil {
		return errors.New("todo.txt is not found")
	}
	defer f.Close()

	r := todotxt.NewReader(f)
	tasks, err := r.ReadAll()
	if err != nil {
		return xerrors.Errorf("Read todo error: %w", err)
	}

	index, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf("args[0] should be int, got %s", c.Args().First())
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("args[0] should be 1-%d", len(tasks))
	}
	tasks[index-1].Reopen()

	f, err = os.Create(todotxtFile)
	if err != nil {
		return errors.New("todo.txt is not found")
	}
	defer f.Close()

	w := todotxt.NewWriter(f)
	err = w.WriteAll(tasks)
	if err != nil {
		return xerrors.Errorf("Failed to write tasks: %w", err)
	}

	fmt.Printf("Succeed to done task(%d)\n", index)
	return nil
}
