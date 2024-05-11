package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}
	DB.Migrator().CreateTable(&Todo{})
	DB.Migrator().CreateTable(&User{})

}

type LocalTime time.Time

// 序列化时间格式
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	// LocalTime 转换成 time.Time 类型
	tTime := time.Time(*t)
	fmt.Println(1)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// 反序列化时间处理
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(string(data), "\"")
	//转为当前服务器所在时区时间
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	// time.Time转换为LocalTime类型
	*t = LocalTime(t1)
	if err != nil {
		return errors.New("时间格式有误，转换失败")
	}
	fmt.Println(2)

	return nil
}
func (t LocalTime) Value() (driver.Value, error) {
	fmt.Println(3)
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixMicro() == zeroTime.UnixMicro() {
		return nil, nil
	}
	return tlt, nil
}
func (t *LocalTime) Scan(v any) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	fmt.Println(4)
	return fmt.Errorf("不能转换 %V 为时间戳", v)
}
