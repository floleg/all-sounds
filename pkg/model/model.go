package model

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
}

type User struct {
	Entity
	UserTrackFavorites []UserTrackFavorite
}

type Album struct {
	Entity
	Title       string
	ReleaseYear uint8
	Tracks      []Track
}

type Track struct {
	Entity
	Title              string
	AlbumID            uint
	Order              uint8
	ArtistID           uint
	UserTrackFavorites []UserTrackFavorite
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
