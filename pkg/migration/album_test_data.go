package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"

	"github.com/bxcodec/faker/v4"
)

func BulkInsertAlbums(number int) []model.Album {
	var albums = []model.Album{}

	for i := 1; i <= number; i++ {
		album := model.Album{}
		faker.FakeData(&album)
		albums = append(albums, album)
	}

	db.DBCon.Create(&albums)

	return albums
}

func InsertAlbum(title string) model.Album {
	var album = model.Album{}

	faker.FakeData(&album)
	album.Title = title

	db.DBCon.Create(&album)

	return album
}
