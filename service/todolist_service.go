package service

type TodolistService interface {
	ShowTodolist()
	Add(todo string)
	Delete(id int64)
}
