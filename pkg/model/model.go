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

type Track struct {
	Entity
	Title    string `faker:"sentence"`
	Order    uint8
	ArtistID uint
	Users    []User  `gorm:"many2many:user_tracks;" faker:"-"`
	Albums   []Album `gorm:"many2many:album_tracks;" faker:"-"`
}

type Artist struct {
	Entity
	Name string
}
