package main

import (
	repository2 "Go_Todolist/repository"
	service2 "Go_Todolist/service"
	"Go_Todolist/util"
	view2 "Go_Todolist/view"
)

func main() {
	repository := repository2.MakeTodolistRepository(util.GetConnection())
	service := service2.MakeTodolistService(repository)
	view := view2.MakeTodolistView(service)
	view.ShowAllTodo()
}
