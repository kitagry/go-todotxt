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
	Flags:  removeFlags,
}

var commandList = &cli.Command{
	Name:   "ls",
	Usage:  "List todo.txt",
	Action: todotxtList,
	Flags:  listFlags,
}

var commandPriority = &cli.Command{
	Name:   "pri",
	Usage:  "Change priority",
	Action: todotxtPriority,
	Flags:  priorityFlags,
}

var commandDo = &cli.Command{
	Name:   "do",
	Usage:  "Complete task",
	Action: todotxtDo,
	Flags:  doFlags,
}

var commandUndo = &cli.Command{
	Name:   "undo",
	Usage:  "Uncomplete task",
	Action: todotxtUndo,
	Flags:  undoFlags,
}
