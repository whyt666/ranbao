package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(225)"`
	Phone    string `gorm:"type:varchar(225)"`
	Email    string `gorm:"type:varchar(225)"`
	Avatar   string `gorm:"type:varchar(225)"`
	LevelNum int64
	SportId  int64
}
