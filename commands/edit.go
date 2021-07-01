package commands

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/jormin/todo/config"
	"github.com/jormin/todo/entity"
	"github.com/jormin/todo/errors"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:      "edit",
			Usage:     "edit todo",
			Action:    Edit,
			ArgsUsage: "[id: unique todo id]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "c",
					Usage:       "the content of todo",
					Required:    false,
					DefaultText: time.Now().Format(""),
				},
				&cli.StringFlag{
					Name:        "d",
					Usage:       "the date of todo",
					Required:    false,
					DefaultText: "today",
				},
				&cli.IntFlag{
					Name:        "l",
					Usage:       "the level of todo, optional values are 1-3, 1 is the minimum level and 3 is the maximum level",
					Required:    false,
					DefaultText: "1",
				},
				&cli.IntFlag{
					Name:        "s",
					Usage:       "the status of todo, optional values are 0-1, 0 means incomplete and 1 means completed",
					Required:    false,
					DefaultText: "0",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// Edit edit todo
func Edit(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	id := ctx.Args().Get(0)
	todo, ok := (*data.Todos)[id]
	if !ok {
		return errors.TodoNotExistsErr
	}
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "c":
			content := ctx.String("c")
			if content == "" {
				return errors.FlagContentValidateErr
			}
			todo.Content = content
		case "d":
			date := ctx.String("d")
			if date == "" {
				return errors.FlagDateValidateErr
			}
			if date == "today" {
				date = time.Now().Format("20060102")
			} else {
				_, err := time.Parse("20060102", date)
				if err != nil {
					return errors.FlagDateValidateErr
				}
			}
			todo.Date = date
		case "l":
			level := ctx.Int("l")
			if level < entity.TodoLevelLow || level > entity.TodoLevelHigh {
				return errors.FlagLevelValidateErr
			}
			todo.Level = level
		case "s":
			status := ctx.Int("s")
			if status < entity.TodoStatusIncomplete || status > entity.TodoStatusCompleted {
				return errors.FlagStatusValidateErr
			}
			todo.Status = status
		}
	}
	(*data.Todos)[id] = todo
	fmt.Printf("edit todo %s success\n", todo.ID)
	return nil
}
