package service

import (
	"errors"
	"todoList-golang/model"
)

func CreateTodoTaskServer(todo *model.Todo) error {
	if todo.TaskType < 1 || todo.TaskType > 4 {
		return errors.New("类型错误" + string(todo.TaskType))
	}
	err := model.CreateTodo(todo)
	return err
}
