package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jormin/todo/entity"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

var data *entity.Data

// Get current data
func GetData() *entity.Data {
	if data == nil {
		data = NewData()
	}
	if data.Todos == nil {
		data.Todos = make(map[string]entity.Todo)
	}
	return data
}

// Get new data
func NewData() *entity.Data {
	return &entity.Data{
		Todos: make(map[string]entity.Todo),
	}
}

// Get path of data file
func getDatafilePath() string {
	home, _ := homedir.Dir()
	return fmt.Sprintf("%s/todo.json", home)
}

// Before
func BeforeFunc(ctx *cli.Context) error {
	path := getDatafilePath()
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err = os.Create(path)
		if err != nil {
			return err
		}
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if string(b) == "" {
		data = NewData()
	} else {
		err = json.Unmarshal(b, &data)
		if err != nil {
			return err
		}
	}
	return nil
}

// After
func AfterFunc(ctx *cli.Context) error {
	path := getDatafilePath()
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0777)
	return err
}
