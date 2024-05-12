package model

type Todo struct {
	ID        uint      `gorm:"primary_key" json:"id"` // 主键ID
	CreatedAt LocalTime `json:"created_at"`            // 创建时间（由GORM自动管理）
	UpdatedAt LocalTime `json:"updated_at"`            // 最后一次更新时间（由GORM自动管理）
	// DeletedAt     gorm.DeletedAt `gorm:"index"`
	Title         string     `json:"title"`          //任务标题
	CompletedTime *LocalTime `json:"completed_time"` // 完成时间
	TaskType      uint       `json:"task_type"`      // 任务类型
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

// 筛选todo
func GetAllTodo(todo *Todo) ([]Todo, error) {
	var todos []Todo
	if err := DB.Find(&todos, todo).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
