package main

import "github.com/urfave/cli/v2"

var commands = []*cli.Command{
	commandAdd,
	commandRemove,
	commandList,
	commandPriority,
	commandDo,
	commandUndo,
}

var defaultFlags = []cli.Flag{
	&cli.StringFlag{Name: "file", Value: "todo.txt", Aliases: []string{"f"}, Usage: "Path to todo.txt file"},
}

var commandAdd = &cli.Command{
	Name:   "add",
	Usage:  "Add task to todo.txt",
	Action: todotxtAdd,
	Flags:  addFlags,
}

var commandRemove = &cli.Command{
	Name:   "rm",
	Usage:  "Remove task from todo.txt",
	Action: todotxtRemove,
	Flags:  defaultFlags,
}

var commandList = &cli.Command{
	Name:   "ls",
	Usage:  "List todo.txt",
	Action: todotxtList,
	Flags:  defaultFlags,
}

var commandPriority = &cli.Command{
	Name:   "pri",
	Usage:  "Change priority",
	Action: todotxtPriority,
	Flags:  defaultFlags,
}

var commandDo = &cli.Command{
	Name:   "do",
	Usage:  "Complete task",
	Action: todotxtDo,
	Flags:  defaultFlags,
}

var commandUndo = &cli.Command{
	Name:   "undo",
	Usage:  "Uncomplete task",
	Action: todotxtUndo,
	Flags:  defaultFlags,
}
