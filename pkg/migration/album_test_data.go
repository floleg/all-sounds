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

func BulkInsertTracksPerArtist(artists []*model.Artist, number int) []*model.Track {
	var tracks = []*model.Track{}

	for _, artist := range artists {
		for i := 1; i <= number; i++ {
			track := model.Track{}
			faker.FakeData(&track)
			track.ArtistID = artist.ID
			tracks = append(tracks, &track)
		}
	}

	db.DBCon.Create(&tracks)

	return tracks
}

func BulkInsertAlbumTracks(albums []*model.Album, tracks []*model.Track) []*model.AlbumTrack {
	var albumTracks = []*model.AlbumTrack{}

	for _, album := range albums {
		for i := 1; i <= 10; i++ {
			albumTrack := model.AlbumTrack{}
			albumTrack.AlbumID = album.ID
			albumTracks = append(albumTracks, &albumTrack)
		}
	}

	db.DBCon.Create(&albumTracks)

	return albumTracks
}

func BulkInsertAlbums(number int) []*model.Album {
	var albums = []*model.Album{}

	for i := 1; i <= number; i++ {
		album := model.Album{}
		faker.FakeData(&album)
		albums = append(albums, &album)
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
