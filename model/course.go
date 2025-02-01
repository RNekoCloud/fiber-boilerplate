package model

import (
	"time"

	"gorm.io/gorm"
)

/*
* Model is simply database enitty for query data.
* Course is a where teacher can add their own study material.
 */
type Course struct {
	ID           string `gorm:"primaryKey"`
	TeacherID    string
	Title        string
	Description  string
	Slug         string
	ThumbnailImg string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Material struct {
	ID        string `gorm:"primaryKey"`
	CourseID  string
	Title     string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SubMaterial struct {
	ID         string `gorm:"primaryKey"`
	MaterialID string
	Title      string
	Type       string `gorm:"type:varchar(20)"`
	Slug       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type StudyMaterial struct {
	ID            string `gorm:"primaryKey"`
	SubMaterialID string
	Content       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
