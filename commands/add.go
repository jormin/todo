package commands

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/urfave/cli/v2"
	"github.com/jormin/todo/config"
	"github.com/jormin/todo/entity"
	"github.com/jormin/todo/errors"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:      "add",
			Usage:     "add todo",
			Action:    Add,
			ArgsUsage: "[content: todo content]",
			Flags: []cli.Flag{
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

// Add add todo
func Add(ctx *cli.Context) error {
	if ctx.Args().Len() == 0 {
		return errors.MissingRequiredArgumentErr
	}
	date := ctx.String("d")
	level := ctx.Int("l")
	status := ctx.Int("s")
	if date != "" {
		if date == "today" {
			date = time.Now().Format("20060102")
		} else {
			_, err := time.Parse("20060102", date)
			if err != nil {
				return errors.FlagDateValidateErr
			}
		}
	}
	if level < entity.TodoLevelLow || level > entity.TodoLevelHigh {
		return errors.FlagLevelValidateErr
	}
	if status < entity.TodoStatusIncomplete || status > entity.TodoStatusCompleted {
		return errors.FlagStatusValidateErr
	}
	content := ctx.Args().Get(0)
	curTime := time.Now().Unix()
	id := xid.New().String()
	(*data.Todos)[id] = entity.Todo{
		ID:         id,
		Content:    content,
		Date:       date,
		Status:     status,
		Level:      level,
		CreateTime: curTime,
		UpdateTime: curTime,
	}
	fmt.Println(id)
	return nil
}
