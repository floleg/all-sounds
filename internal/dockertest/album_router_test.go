package dockertest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"allsounds/internal/router"
	"allsounds/pkg/migration"
	"allsounds/pkg/model"

	"github.com/magiconair/properties/assert"
)

// Without offset or limit url parameters, endpoint will return 400
func TestFindAllWithoutPagination(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
	assert.Equal(t, w.Body.String(), "{\"message\":\"bad request\"}")
}

// Test endpoint pagination
func TestFindAll(t *testing.T) {
	migration.BulkInsertAlbums(1000)

	router := router.NewRouter()

	w := httptest.NewRecorder()

	for i := 0; i <= 900; i += 100 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/album?offset=%v&limit=100", i), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, 200)

		data := []model.Album{}
		json.NewDecoder(w.Body).Decode(&data)
		assert.Equal(t, len(data), 100)
	}
}

// Passing a string as ID should return 400
func TestAlbumByIdWithString(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/album/misguided-id", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
	assert.Equal(t, w.Body.String(), "{\"message\":\"bad request\"}")
}

func TestAlbumById(t *testing.T) {
	album := migration.InsertAlbum("Album Title 1")

	router := router.NewRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", fmt.Sprintf("/album/%v", album.ID), nil)
	router.ServeHTTP(w, req)

	data := model.Album{}

	assert.Equal(t, w.Code, 200)
	json.NewDecoder(w.Body).Decode(&data)

	assert.Equal(t, &album.ID, &data.ID)
	assert.Equal(t, &album.Title, &data.Title)
}

func TestSearch(t *testing.T) {
	router := router.NewRouter()

	for i := 0; i <= 50; i++ {
		migration.InsertAlbum(fmt.Sprintf("Album Title %v", i))
	}

	query := "Title"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/album?query=%s&offset=0&limit=100", query), nil)
	router.ServeHTTP(w, req)

	data := []model.Album{}
	json.NewDecoder(w.Body).Decode(&data)
	assert.Equal(t, len(data), 51)
}
