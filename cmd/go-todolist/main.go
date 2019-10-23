package main

import "github.com/danteay/go-todolist/common/models"

func main() {
	list := new(models.TodoList)
	println(list.Name)
}
