package models

import (
	"time"
)

type User struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"size:255;not null"`
	Email     string     `json:"email" gorm:"size:255;not null;unique"`
	Phone     string     `json:"phone" gorm:"size:17"`
	Password  string     `json:"password" gorm:"size:255"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

type UserIndex struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"size:255;not null"`
	Email     string     `json:"email" gorm:"size:255;not null;unique"`
	Phone     string     `json:"phone" gorm:"size:17"`
	Password  string     `json:"password" gorm:"size:255"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

func (UserIndex) TableName() string {
	return "public.user"
}

type UserExample struct {
	Name     string `json:"name" gorm:"size:255;not null"`
	Email    string `json:"email" gorm:"size:255;not null;unique"`
	Phone    string `json:"phone" gorm:"size:17"`
	Password string `json:"password" gorm:"size:255"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
