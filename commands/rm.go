package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/jormin/todo/config"
	"github.com/jormin/todo/errors"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:      "rm",
			Usage:     "remove todo",
			Action:    Remove,
			ArgsUsage: "[id1: unique todo id] [id2] ... [idn]",
			Before:    BeforeFunc,
			After:     AfterFunc,
		},
	)
}

// Remove remove todo
func Remove(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	for i := 0; i < ctx.Args().Len(); i++ {
		id := ctx.Args().Get(i)
		for index, item := range *data.Todos {
			if item.ID == id {
				delete(*data.Todos, index)
				fmt.Printf("remove todo %s success\n", item.ID)
			}
		}
	}
	return nil
}
