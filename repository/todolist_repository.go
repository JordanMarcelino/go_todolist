package repository

import (
	"Go_Todolist/entity"
	"context"
)

type TodolistRepository interface {
	Insert(ctx context.Context, todolist entity.Todolist) (entity.Todolist, error)
	Delete(ctx context.Context, id int64) (bool, error)
	GetAllTodo(ctx context.Context) ([]entity.Todolist, error)
}
