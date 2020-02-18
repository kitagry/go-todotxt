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

var priorityFlags = []cli.Flag{
	&cli.StringFlag{Name: "file", Value: "todo.txt", Aliases: []string{"f"}, Usage: "Path to todo.txt file"},
}

func todotxtPriority(c *cli.Context) error {
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
	index--

	pris := []byte(c.Args().Get(1))
	if len(pris) != 1 {
		return fmt.Errorf("args[1] should be A-Z, got %s", c.Args().Get(1))
	}
	err = tasks[index].SetPriority(pris[0])
	if err != nil {
		return xerrors.Errorf("Priority update error: %w", err)
	}

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

	fmt.Printf("Succeed to change priority to %s", string(tasks[index].Priority()))
	return nil
}
