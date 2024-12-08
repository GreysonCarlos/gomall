package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex;type:varchar(128) not null"`
	Password string  `gorm:"type:varchar(64) not null"`
}

// 防止表的名字被命名为蛇形负数一样,可以使用如下函数进行表的命名
func (User) TableName() string {
	return "user"
}