package model

import "allsounds/pkg/db"

type Album struct {
	Entity
	Title       string
	ReleaseYear uint8
	Tracks      []Track
}

func (a Album) FindAll(offset int, limit int) []Album {
	var albums []Album
	db.DBCon.Order("title asc").Limit(limit).Offset(offset).Find(&albums)

	return albums
}
