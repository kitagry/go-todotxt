package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kitagry/go-todotxt"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var addFlags = []cli.Flag{
	&cli.StringFlag{Name: "file", Value: "todo.txt", Aliases: []string{"f"}, Usage: "Path to todo.txt file"},
	&cli.StringFlag{Name: "pri", Value: "", Aliases: []string{"p"}, Usage: "Set priority to task"},
}

func todotxtAdd(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("args length should be 1, got %v", c.Args().Slice())
	}
	todotxtFile := c.String("file")
	f, err := os.OpenFile(todotxtFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return errors.New("todo.txt is not found")
	}
	defer f.Close()

	w := todotxt.NewWriter(f)
	task := todotxt.NewTask()
	if p := []byte(c.String("pri")); len(p) != 0 {
		if len(p) > 1 {
			return fmt.Errorf("Priority should be A-Z, got %s", string(p))
		}

		err := task.SetPriority(p[0])
		if err != nil {
			return xerrors.Errorf("SetPriority(%s) error: %w", p, err)
		}
	}
	task.SetDescription(c.Args().First())
	err = w.Write(task)
	if err != nil {
		return xerrors.Errorf("Failed to write to %s: %w", todotxtFile, err)
	}
	err = w.Flush()
	if err != nil {
		return xerrors.Errorf("Failed to flush to %s: %w", todotxtFile, err)
	}

	return nil
}
