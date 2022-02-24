package repository

import (
	"Go_Todolist/entity"
	"context"
	"database/sql"
	"errors"
)

type TodolistRepositoryImpl struct {
	Db *sql.DB
}

func MakeTodolistRepository(db *sql.DB) TodolistRepository {
	return &TodolistRepositoryImpl{
		Db: db,
	}
}

func (t *TodolistRepositoryImpl) Insert(ctx context.Context, todolist entity.Todolist) (entity.Todolist, error) {
	query, err := t.Db.PrepareContext(ctx, "INSERT INTO todolist(todo) VALUES(?)")
	if err != nil {
		return todolist, err
	}

	result, execError := query.ExecContext(ctx, todolist.Todo)
	if execError != nil {
		return todolist, execError
	}

	id, idErr := result.LastInsertId()
	if idErr != nil {
		return todolist, err
	}

	todolist.Id = id
	return todolist, nil
}

func (t *TodolistRepositoryImpl) isExist(ctx context.Context, id int64) bool {
	query, err := t.Db.PrepareContext(ctx, "SELECT id FROM todolist WHERE id = ?")
	if err != nil {
		return false
	}

	res, execError := query.QueryContext(ctx, id)
	if execError != nil {
		return false
	}

	if res.Next() {
		return true
	}
	return false
}

func (t *TodolistRepositoryImpl) Delete(ctx context.Context, id int64) (bool, error) {
	if t.isExist(ctx, id) {
		query, err := t.Db.PrepareContext(ctx, "DELETE FROM todolist WHERE id = ?")
		if err != nil {
			return false, errors.New("todolist not found")
		}

		_, execError := query.ExecContext(ctx, id)
		if execError != nil {
			return false, execError
		}

		return true, nil
	}
	return false, errors.New("todolist not found")
}

func (t *TodolistRepositoryImpl) GetAllTodo(ctx context.Context) ([]entity.Todolist, error) {
	query, err := t.Db.PrepareContext(ctx, "SELECT id, todo FROM todolist")
	if err != nil {
		return nil, err
	}

	result, execError := query.QueryContext(ctx)
	if execError != nil {
		return nil, execError
	}

	var listTodo []entity.Todolist
	temp := entity.Todolist{}
	for result.Next() {
		err := result.Scan(&temp.Id, &temp.Todo)
		if err != nil {
			return nil, err
		}
		listTodo = append(listTodo, temp)
	}

	return listTodo, nil
}
