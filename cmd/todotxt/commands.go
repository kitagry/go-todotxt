package main

import "github.com/urfave/cli/v2"

var commands = []*cli.Command{
	commandList,
	commandAdd,
}

var commandAdd = &cli.Command{
	Name:   "add",
	Usage:  "Add task to todo.txt",
	Action: todotxtAdd,
	Flags:  addFlags,
}

var commandList = &cli.Command{
	Name:   "ls",
	Usage:  "List todo.txt",
	Action: todotxtList,
	Flags:  listFlags,
}
