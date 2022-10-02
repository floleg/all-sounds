package model

import (
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	ID        uint `gorm:"primaryKey" faker:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}

type User struct {
	Entity
	Tracks []Track `gorm:"many2many:user_tracks;" faker:"-"`
}
