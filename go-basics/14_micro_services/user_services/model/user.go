package model

type User struct {
	GormModel
	ID           int32  `gorm:"primary_key"`
	MobileNumber string `gorm:"index:idx_mobile_number;unique;type:varchar(11);not null"`
	NickName     string `gorm:"type:varchar(15);unique"`
	Password     string `gorm:"type:varchar(100);not null"`
	Gender       int32  `gorm:"default:0;type:tinyint(1) comment '0 for male, 1 for female'"`
	Role         int32  `gorm:"default:0;type:tinyint(1) comment '0 for banned-user, 1 for common-user, 2 for admin'"`
}
