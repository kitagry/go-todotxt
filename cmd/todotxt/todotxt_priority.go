package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

func todotxtPriority(c *cli.Context) error {
	todotxtFile := c.String("file")
	tasks, err := getTasks(todotxtFile)
	if err != nil {
		return xerrors.Errorf("Failed to getTasks: %w", err)
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

	err = overWrite(todotxtFile, tasks)
	if err != nil {
		return xerrors.Errorf("Failed to write tasks: %w", err)
	}

	fmt.Printf("Succeed to change priority to %s\n", string(tasks[index].Priority()))
	return nil
}
