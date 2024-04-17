package model

import (
	"database/sql"
	"time"
)

type gormModel struct {
	ID        uint         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deletedAt" gorm:"index"`
}
