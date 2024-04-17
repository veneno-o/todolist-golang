package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Create(todo *Todo) error {
	var err error
	if err = DB.Create(todo).Error; err != nil {
		return nil
	}
	return err
}

func Delete(id string) error {
	var err error
	if err = DB.Delete(&Todo{}, id).Error; err != nil {
		return nil
	}
	return err
}

func Update(todo *Todo) error {
	var err error
	if err = DB.Save(todo).Error; err != nil {
		return nil
	}
	return err
}

func Get(todo *Todo) (*Todo, error) {
	var err error
	if err = DB.Find(&todo).Error; err != nil {
		return todo, nil
	}
	return nil, err
}
