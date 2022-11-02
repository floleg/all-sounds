package migration

import (
	"allsounds/pkg/db"
	"allsounds/pkg/model"
	"github.com/bxcodec/faker/v4"
	"github.com/rs/zerolog/log"
)

// BulkInsertArtists inserts a set of artists in the artist table
// depending on the given number parameter.
// Returns a slice of Artist.
func BulkInsertArtists(number int) []*model.Artist {
	var artists []*model.Artist

	for i := 1; i <= number; i++ {
		artist := model.Artist{}
		err := faker.FakeData(&artist)
		if err != nil {
			log.Err(err)
		}
		artists = append(artists, &artist)
	}

	db.DBCon.Create(&artists)

	return artists
}

// BulkInsertAlbums inserts a set of albums in the `album` table
// depending on the given number parameter.
// Returns a slice of Album.
func BulkInsertAlbums(artists []*model.Artist, number int) []*model.Album {
	var albums []*model.Album

	for _, artist := range artists {
		for i := 1; i <= number; i++ {
			album := model.Album{}
			err := faker.FakeData(&album)
			if err != nil {
				log.Err(err)
			}

			for i := 1; i <= number; i++ {
				track := model.Track{}
				err = faker.FakeData(&track)
				if err != nil {
					log.Err(err)
				}

				track.ArtistID = artist.ID
				album.Tracks = append(album.Tracks, track)
			}

			albums = append(albums, &album)
		}
	}

	db.DBCon.Create(&albums)

	return albums
}

// BulkInsertUsers inserts a set of users in the `user` table
// depending on the given number parameter.
// Returns a slice of User.
func BulkInsertUsers(number int) []model.User {
	var users []model.User

	for i := 0; i < number; i++ {
		user := model.User{}
		err := faker.FakeData(&user)
		if err != nil {
			log.Err(err)
		}
		users = append(users, user)
	}

	db.DBCon.Create(&users)

	return users
}
