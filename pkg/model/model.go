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
	UserTrackFavorites []UserTrackFavorite
}

type Track struct {
	Entity
	Title              string `faker:"sentence"`
	Order              uint8
	ArtistID           uint
	UserTrackFavorites []UserTrackFavorite
	AlbumTracks        []AlbumTrack `faker:"-"`
}

type AlbumTrack struct {
	Entity
	AlbumID uint
	TrackID uint
}

type Artist struct {
	Entity
	Name   string
	Tracks []Track
}

type UserTrackFavorite struct {
	Entity
	UserID  uint
	TrackID uint
}
