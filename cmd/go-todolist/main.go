package main

import (
	"github.com/danteay/go-todolist/common/database"
	"github.com/danteay/go-todolist/common/models"
	"github.com/danteay/go-todolist/common/serialization"
)

func main() {
	list := new(models.TodoList)
	connp := new(database.Postgres)
	connm := new(database.Mysql)

	list.Name = "some"
	list.Description = "asdasd"

	list.Store(connp)
	list.Store(connm)

	data := serialization.JSON{
		"param": "asdasd",
		"var":   12312,
	}
}
