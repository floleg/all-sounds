package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"

	"github.com/bxcodec/faker/v4"
)

func BulkInsertAlbums() {
	var albums = []model.Album{}

	for i := 1; i <= 1000; i++ {
		album := model.Album{}
		faker.FakeData(&album)
		albums = append(albums, album)
	}

	db.DBCon.Create(&albums)
}
