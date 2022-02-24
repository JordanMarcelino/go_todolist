package test

import (
	"Go_Todolist/entity"
	repository2 "Go_Todolist/repository"
	"Go_Todolist/util"
	"context"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	repository := repository2.MakeTodolistRepository(util.GetConnection())
	ctx := context.Background()
	todolist := entity.Todolist{
		Todo: "Android Jetpack",
	}
	_, err := repository.Insert(ctx, todolist)
	if err != nil {
		panic(err)
	}
}

func TestDelete(t *testing.T) {
	repository := repository2.MakeTodolistRepository(util.GetConnection())
	ctx := context.Background()

	_, err := repository.Delete(ctx, 4)
	if err != nil {
		panic(err)
	}
}

func TestGetAll(t *testing.T) {
	repository := repository2.MakeTodolistRepository(util.GetConnection())
	ctx := context.Background()
	res, err := repository.GetAllTodo(ctx)
	if err != nil {
		panic(err)
	}

	for _, re := range res {
		fmt.Println(re)
	}
}
