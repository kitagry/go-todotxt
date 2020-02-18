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

var removeFlags = []cli.Flag{
	&cli.StringFlag{Name: "file", Value: "todo.txt", Aliases: []string{"f"}, Usage: "Path to todo.txt file"},
}

func todotxtRemove(c *cli.Context) error {
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
		return fmt.Errorf("args should be int, got %s", c.Args().First())
	}
	tasks = removeTask(tasks, index-1)

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
