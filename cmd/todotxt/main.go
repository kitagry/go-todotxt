package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "todotxt"
	app.Usage = "Manage todo.txt"
	app.Authors = []*cli.Author{
		{
			Name:  "kitagry",
			Email: "kitadrum50@gmail.com",
		},
	}
	app.Commands = commands
	return app
}

func main() {
	err := newApp().Run(os.Args)
	if err != nil {
		exitCode := 1
		if excoder, ok := err.(cli.ExitCoder); ok {
			exitCode = excoder.ExitCode()
		}
		fmt.Println(err)
		os.Exit(exitCode)
	}
}
