package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

func todotxtUndo(c *cli.Context) error {
	todotxtFile := c.String("file")
	tasks, err := getTasks(todotxtFile)
	if err != nil {
		return xerrors.Errorf("Failed to getTasks: %w", err)
	}

	index, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf("args[0] should be int, got %s", c.Args().First())
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("args[0] should be 1-%d", len(tasks))
	}
	tasks[index-1].Reopen()

	err = overWrite(todotxtFile, tasks)
	if err != nil {
		return xerrors.Errorf("Failed to overWrite: %w", err)
	}

	fmt.Printf("Succeed to done task(%d)\n", index)
	return nil
}
