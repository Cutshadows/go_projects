package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null;unique_index"`
	Description string
	Done        bool `gorm:"default:false"`
	UserID      uint
}
