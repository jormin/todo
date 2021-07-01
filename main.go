package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	_ "github.com/jormin/todo/commands"
	"github.com/jormin/todo/config"
)

func main() {
	app := &cli.App{
		Name:        "todo",
		Usage:       "A simple tool to manage your todo list",
		Version:     "v1.0.0",
		Description: "A simple tool to manage your todo list",
		Commands:    config.GetRegisteredCommands(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
