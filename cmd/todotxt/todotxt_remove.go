package main

import (
	"fmt"
	"strconv"

	"github.com/kitagry/go-todotxt"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

func todotxtRemove(c *cli.Context) error {
	todotxtFile := c.String("file")
	tasks, err := getTasks(todotxtFile)
	if err != nil {
		return xerrors.Errorf("Failed to getTasks: %w", err)
	}

	index, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf("args should be int, got %s", c.Args().First())
	}
	tasks = removeTask(tasks, index-1)

	err = overWrite(todotxtFile, tasks)
	if err != nil {
		return xerrors.Errorf("Failed to write tasks: %w", err)
	}

	fmt.Printf("Succeed to remove task(%d)\n", index)
	return nil
}

func removeTask(list []*todotxt.Task, index int) []*todotxt.Task {
	newList := make([]*todotxt.Task, len(list))
	copy(newList, list)
	if index == len(newList)-1 {
		return newList[:len(newList)-1]
	} else if index == 0 {
		return newList[1:]
	} else if index >= len(newList) || index < 0 {
		return newList
	}
	return append(newList[:index], newList[index+1:]...)
}
