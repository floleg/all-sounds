package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"

	"github.com/bxcodec/faker/v4"
)

func BulkInsertArtists(number int) []*model.Artist {
	var artists = []*model.Artist{}

	for i := 1; i <= number; i++ {
		artist := model.Artist{}
		faker.FakeData(&artist)
		artists = append(artists, &artist)
	}

	db.DBCon.Create(&artists)

	return artists
}

func BulkInsertAlbums(artists []*model.Artist, number int) []*model.Album {
	var albums = []*model.Album{}

	for _, artist := range artists {
		for i := 1; i <= number; i++ {
			album := model.Album{}
			faker.FakeData(&album)

			for i := 1; i <= number; i++ {
				track := model.Track{}
				faker.FakeData(&track)
				track.ArtistID = artist.ID
				album.Tracks = append(album.Tracks, track)
			}

			albums = append(albums, &album)
		}
	}

	db.DBCon.Create(&albums)

	return albums
}

func BulkInsertUsers(number int) []model.User {
	var users = []model.User{}

	for i := 0; i < number; i++ {
		user := model.User{}
		faker.FakeData(&user)
		users = append(users, user)
	}

	db.DBCon.Create(&users)

	return users
}
