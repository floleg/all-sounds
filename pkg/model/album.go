package model

import (
	"allsounds/pkg/db"
)

type Album struct {
	Entity
	Title       string `faker:"sentence"`
	ReleaseYear uint8
	Tracks      []Track `faker:"-"`
}

func (a Album) FindAll(offset int, limit int) []Album {
	var albums []Album
	db.DBCon.Order("title asc").Limit(limit).Offset(offset).Find(&albums)

	return albums
}

func (a Album) Search(offset int, limit int, query string) []Album {
	var albums []Album
	db.DBCon.Order("title asc").Limit(limit).Offset(offset).Where("title LIKE ?", "%"+query+"%").Find(&albums)

	return albums
}

func (a Album) FindById(id int) Album {
	var album Album
	db.DBCon.First(&album, id)

	return album
}
