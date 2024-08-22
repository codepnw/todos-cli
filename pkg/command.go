package pkg

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codepnw/todo-cli/types"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & new title [ex. id:new_title]")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo to completed")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()
	return &cf
}

func (cf *cmdFlags) Execute(todos *types.Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format edit; please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index for edit")
			os.Exit(1)
		}
		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Del != -1:
		todos.Delete(cf.Del)

	default:
		fmt.Println("invalid command")
	}
}
