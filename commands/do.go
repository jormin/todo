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
			Name:      "do",
			Usage:     "do or undo(with -r) todo",
			Action:    Do,
			ArgsUsage: "[id1: unique todo id] [id2] ... [idn]",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "a",
					Usage:       "do or undo(with -r) all todo",
					Required:    false,
					DefaultText: "false",
				},
				&cli.BoolFlag{
					Name:        "r",
					Usage:       "undo todo",
					Required:    false,
					DefaultText: "false",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// Do do or undo(with -r) todo
func Do(ctx *cli.Context) error {
	doAll := false
	undo := false
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			doAll = ctx.Bool("a")
		case "r":
			undo = ctx.Bool("r")
		}
	}
	if !doAll && ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	label := "do"
	status := entity.TodoStatusCompleted
	if undo {
		label = "undo"
		status = entity.TodoStatusUncompleted
	}
	for index, item := range data.Todos {
		if doAll {
			item.Status = status
			data.Todos[index] = item
			fmt.Printf("%s todo %s success\n", label, item.ID)
		} else {
			for i := 0; i < ctx.Args().Len(); i++ {
				id := ctx.Args().Get(i)
				if item.ID == id {
					item.Status = status
					data.Todos[index] = item
					fmt.Printf("%s todo %s success\n", label, item.ID)
				}
			}
		}
	}
	return nil
}
