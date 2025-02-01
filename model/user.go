package model

import (
	"time"

	"gorm.io/gorm"
)

type ROLE string

const (
	SUPER_ADMIN ROLE = "SUPER_ADMIN"
	ADMIN       ROLE = "ADMIN"
	TEACHER     ROLE = "TEACHER"
	STUDENT     ROLE = "STUDENT"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string
	Password  string
	Name      string
	Role      ROLE
	CreateAt  time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Teacher struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
