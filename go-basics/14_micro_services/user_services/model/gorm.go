package model

import (
	"time"

	"gorm.io/gorm"
)

// gorm 提供的预定义的结构体

type GormModel struct {
	CreateAt  *time.Time      `json:"create_at" gorm:"default: null"`
	UpdateAt  *time.Time      `json:"update_at" gorm:"default: null"`
	DeleteAt  *gorm.DeletedAt `json:"delete_at" gorm:"index;default: null"`
	IsDeleted bool            `json:"is_deleted"`
}
