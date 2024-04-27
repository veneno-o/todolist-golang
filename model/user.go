package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name     string `json:"name"`
	password string `json:"password"`
}

// 增
func CreateUser(todo *Todo) error {
	var err error
	if err = DB.Create(todo).Error; err != nil {
		return nil
	}
	return err
}

// 删
func DeleteUser(id int) error {
	var err error
	if err = DB.Delete(&Todo{}, id).Error; err != nil {
		return nil
	}
	return err
}

// 改
func UpdateUser(todo *Todo) error {
	var err error
	if err = DB.Save(todo).Error; err != nil {
		return nil
	}
	return err
}

// 查
func GetUser(id int) (*Todo, error) {
	var todo Todo
	if err := DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// 查全部
func GetUsers() ([]Todo, error) {
	var todos []Todo
	if err := DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
