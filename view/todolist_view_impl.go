package view

import (
	"Go_Todolist/service"
	"fmt"
	"os"
	"strconv"
)

type TodolistViewImpl struct {
	Service service.TodolistService
}

func MakeTodolistView(service service.TodolistService) TodolistView {
	return &TodolistViewImpl{
		Service: service,
	}
}

func (t *TodolistViewImpl) ShowAllTodo() {
	for {
		fmt.Println("TODOLIST")
		t.Service.ShowTodolist()
		fmt.Println("MENU")
		fmt.Println("1. Add")
		fmt.Println("2. Delete")
		fmt.Println("x. Exit")

		var choose string
		_, err := fmt.Scan(&choose)
		if err != nil {
			panic(err)
		}

		switch choose {
		case "1":
			t.AddTodo()
		case "2":
			t.DeleteTodo()
		case "x":
			os.Exit(0)
		default:
			fmt.Println("Unknown input")
		}
	}
}

func (t *TodolistViewImpl) AddTodo() {
	var input string
	fmt.Print("Enter todolist (x to cancel) : ")
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}

	if input == "x" {
		t.ShowAllTodo()
	}
	t.Service.Add(input)
	t.ShowAllTodo()
}

func (t *TodolistViewImpl) DeleteTodo() {
	var input string
	fmt.Print("Enter todolist id you want to delete (x to cancel) : ")
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	if input == "x" {
		t.ShowAllTodo()
	}

	id, convErr := strconv.Atoi(input)
	if convErr != nil {
		panic(convErr)
	}

	t.Service.Delete(int64(id))
	t.ShowAllTodo()
}
