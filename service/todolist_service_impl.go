package service

import (
	"Go_Todolist/entity"
	"Go_Todolist/repository"
	"context"
	"fmt"
	"strconv"
)

type TodolistServiceImpl struct {
	Repository repository.TodolistRepository
}

func MakeTodolistService(todolistRepository repository.TodolistRepository) TodolistService {
	return &TodolistServiceImpl{
		Repository: todolistRepository,
	}
}

func (t *TodolistServiceImpl) ShowTodolist() {
	ctx := context.Background()
	result, err := t.Repository.GetAllTodo(ctx)
	if err != nil {
		panic("Failed getting todolist")
	}
	for _, todolist := range result {
		fmt.Println(strconv.Itoa(int(todolist.Id)) + ". " + todolist.Todo)
	}
}

func (t *TodolistServiceImpl) Add(todo string) {
	ctx := context.Background()
	todolist := entity.Todolist{Todo: todo}
	_, err := t.Repository.Insert(ctx, todolist)
	if err != nil {
		panic("Failed to add todolist")
	}
	fmt.Println("Success adding todolist")
}

func (t *TodolistServiceImpl) Delete(id int64) {
	ctx := context.Background()
	res, err := t.Repository.Delete(ctx, id)
	if err != nil {
		panic("Failed to remove todolist")
	}
	if res {
		fmt.Println("Success deleting todolist")
	} else {
		fmt.Println("Failed deleting todolist")
	}
}
