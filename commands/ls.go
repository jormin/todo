package commands

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/jormin/todo/config"
	"github.com/jormin/todo/entity"
	"github.com/jormin/todo/errors"
	"github.com/urfave/cli/v2"
)

// init
func init() {
	config.RegisterCommand(
		"", &cli.Command{
			Name:      "ls",
			Usage:     "show config fund list",
			Action:    List,
			ArgsUsage: "",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "a",
					Usage:       "show all todo list",
					Required:    false,
					DefaultText: "false",
				},
				&cli.StringFlag{
					Name:        "d",
					Usage:       "show todo list of one date",
					Required:    false,
					DefaultText: "today",
				},
			},
			Before: BeforeFunc,
			After:  AfterFunc,
		},
	)
}

// List show config fund list
func List(ctx *cli.Context) error {
	showAll := false
	curDate := time.Now().Format("20060102")
	date := curDate

	flags := ctx.FlagNames()
	for _, v := range flags {
		switch v {
		case "a":
			showAll = ctx.Bool("a")
		case "d":
			date = ctx.String("d")
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
		}
	}

	// deal todos
	var todos []entity.Todo
	for _, item := range *data.Todos {
		if !showAll && item.Date != date {
			continue
		}
		todos = append(todos, item)
	}
	for i := 0; i < len(todos)-1; i++ {
		for j := 0; j < len(todos)-i-1; j++ {
			if todos[j+1].Status < todos[j].Status {
				todos[j+1], todos[j] = todos[j], todos[j+1]
			}
			if todos[j+1].Status == todos[j].Status {
				if todos[j+1].Level > todos[j].Level {
					todos[j+1], todos[j] = todos[j], todos[j+1]
				}
				if todos[j+1].Level == todos[j].Level {
					if todos[j+1].Date > todos[j].Date {
						todos[j+1], todos[j] = todos[j], todos[j+1]
					}
				}
			}
		}
	}

	// content format
	contentFormat := "%v\t%s\t%s\t%s\t%s\t%s"
	// default format
	defaultFormat := "\033[1;38;38m%s\033[0m\n"
	// current format
	currentFormat := "\033[1;32;32m%s\033[0m\n"
	// important format
	importantFormat := "\033[1;31;31m%s\033[0m\n"
	// completed format
	completedFormat := "\033[1;37;37m%s\033[0m\n"
	headers := []interface{}{"#", "ID", "Date", "Level", "Status", "Content"}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintf(w, defaultFormat, fmt.Sprintf(contentFormat, headers...))
	for index, item := range todos {
		format := defaultFormat
		if item.Status == entity.TodoStatusCompleted {
			format = completedFormat
		} else if item.Level == entity.TodoLevelHigh {
			format = importantFormat
		} else if item.Date == curDate {
			format = currentFormat
		}
		if showAll && item.Date == curDate {
			item.Date = fmt.Sprintf("%s(today)", item.Date)
		}
		str := fmt.Sprintf(
			contentFormat, index+1, item.ID, item.Date, entity.TodoLevelTexts[item.Level],
			entity.TodoStatusTexts[item.Status],
			item.Content,
		)
		_, _ = fmt.Fprintf(w, format, str)
	}
	_ = w.Flush()
	return nil
}
