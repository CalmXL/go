package model

import (
	"gorm.io/gorm"
)

type Student struct {
	// gorm.Model
	gormModel
	Name string `json:"name" gorm:"size:10;default:游客;"`
	Age  uint   `json:"age" gorm:"not null"`
}

func (s *Student) BeforeSave(db *gorm.DB) (err error) {
	// fmt.Println("Before Save")
	return err
}

func (s *Student) BeforeCreate(db *gorm.DB) (err error) {
	// fmt.Println("Before Create")
	return err
}

func (s *Student) AfterCreate(db *gorm.DB) (err error) {
	// fmt.Println("After Create")
	return err
}

func (s *Student) AfterSave(db *gorm.DB) (err error) {
	// fmt.Println("After Save")
	return err
}
