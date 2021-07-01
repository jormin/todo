package commands

import (
	"fmt"
	"time"

	"github.com/jormin/todo/config"
	"github.com/jormin/todo/entity"
	"github.com/jormin/todo/errors"
	"github.com/rs/xid"
	"github.com/urfave/cli/v2"
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
	content := ctx.Args().Get(0)
	curDate := time.Now().Format("20060102")
	date := curDate
	status := entity.TodoStatusIncomplete
	level := entity.TodoLevelLow

	if content == "" {
		return errors.FlagContentValidateErr
	}
	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "d":
			date = ctx.String("d")
			if date == "" {
				return errors.FlagDateValidateErr
			}
			if date == "today" {
				date = curDate
			} else {
				_, err := time.Parse("20060102", date)
				if err != nil {
					return errors.FlagDateValidateErr
				}
			}
		case "l":
			level = ctx.Int("l")
			if level < entity.TodoLevelLow || level > entity.TodoLevelHigh {
				return errors.FlagLevelValidateErr
			}
		case "s":
			status = ctx.Int("s")
			if status < entity.TodoStatusIncomplete || status > entity.TodoStatusCompleted {
				return errors.FlagStatusValidateErr
			}
		}
	}
	curTime := time.Now().Unix()
	id := xid.New().String()
	todo := entity.Todo{
		ID:         id,
		Content:    content,
		Date:       date,
		Status:     status,
		Level:      level,
		CreateTime: curTime,
		UpdateTime: curTime,
	}
	(*data.Todos)[id] = todo
	fmt.Println(id)
	return nil
}
