package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CreateTodo(todo *Todo) error {
	var err error
	if err = DB.Create(todo).Error; err != nil {
		return nil
	}
	return err
}

func DeleteTodo(id string) error {
	var err error
	if err = DB.Delete(&Todo{}, id).Error; err != nil {
		return nil
	}
	return err
}

func UpdateTodo(id string, todo *Todo) error {
	var err error
	if err = DB.Model(&Todo{}).Where("id = ?", id).Updates(todo).Error; err != nil {
		return nil
	}
	return err
}

func GetTodo(id string) (*Todo, error) {
	var todo Todo
	if err := DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func GetAllTodo() ([]Todo, error) {
	var todos []Todo
	if err := DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
