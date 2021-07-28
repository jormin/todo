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
			Name:      "discard",
			Usage:     "discard or reverse discard(with -r) todo",
			Action:    Discard,
			ArgsUsage: "[id1: unique todo id] [id2] ... [idn]",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "a",
					Usage:       "discard or reverse discard(with -r) all todo",
					Required:    false,
					DefaultText: "false",
				},
				&cli.BoolFlag{
					Name:        "r",
					Usage:       "reverse discard todo",
					Required:    false,
					DefaultText: "false",
				},
				&cli.IntFlag{
					Name:        "s",
					Usage:       "the status to reverse, optional values are 0-1[uncompleted|completed]",
					Required:    false,
					DefaultText: "0",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// Discard discard or reverse discard(with -r) todo
func Discard(ctx *cli.Context) error {
	discardAll := false
	reverseDiscard := false
	reverseStatus := 0
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			discardAll = ctx.Bool("a")
		case "r":
			reverseDiscard = ctx.Bool("r")
		case "s":
			reverseStatus = ctx.Int("s")
		}
	}
	if !discardAll && ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	if reverseDiscard && reverseStatus < entity.TodoStatusUncompleted || reverseStatus > entity.TodoStatusCompleted {
		return errors.FlagStatusValidateErr
	}
	label := "discard"
	status := entity.TodoStatusDiscarded
	if reverseDiscard {
		label = "reverse discard"
		status = reverseStatus
	}
	for index, item := range data.Todos {
		if discardAll {
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
