package commands

import (
	"fmt"

	"github.com/jormin/todo/config"
	"github.com/jormin/todo/entity"
	"github.com/jormin/todo/errors"
	"github.com/urfave/cli/v2"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:      "rm",
			Usage:     "remove todo",
			Action:    Remove,
			ArgsUsage: "[id1: unique todo id] [id2] ... [idn]",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "a",
					Usage:       "remove all todo",
					Required:    false,
					DefaultText: "false",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// Remove remove todo
func Remove(ctx *cli.Context) error {
	removeAll := false
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			removeAll = ctx.Bool("a")
		}
	}
	if removeAll {
		*data.Todos = map[string]entity.Todo{}
		fmt.Println("remove all todos success")
	} else {
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
	}
	return nil
}
