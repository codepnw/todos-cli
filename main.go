package main

import (
	"github.com/codepnw/todo-cli/pkg"
	"github.com/codepnw/todo-cli/store"
	"github.com/codepnw/todo-cli/types"
)

func main() {
	todos := types.Todos{}
	storage := store.NewStorage[types.Todos]("todos.json")
	cmdFlags := pkg.NewCmdFlags()

	storage.Load(&todos)
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
