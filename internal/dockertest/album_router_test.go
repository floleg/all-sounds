package dockertest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"allsounds/internal/router"
	"allsounds/pkg/model"
)

// Without offset or limit url parameters, endpoint will return 400
func TestFindAllWithoutPagination(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album", nil)
	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("got %v, want %v", w.Code, 400)
	}

	if w.Body.String() != "{\"message\":\"bad request\"}" {
		t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
	}
}

// Test endpoint pagination
func TestFindAll(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()

	for i := 0; i <= 9; i += 10 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/album?offset=%v&limit=10", i), nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		data := []model.Album{}
		json.NewDecoder(w.Body).Decode(&data)

		if len(data) != 10 {
			t.Errorf("got %v, want %v", len(data), 10)
		}
	}
}

// Passing a string as ID should return 400
func TestAlbumByIdWithString(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/album/misguided-id", nil)
	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("got %v, want %v", w.Code, 400)
	}

	if w.Body.String() != "{\"message\":\"bad request\"}" {
		t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
	}
}

func TestAlbumById(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()

	// First search in album list to retreieve an actual album id
	findAllReq, _ := http.NewRequest("GET", "/album?offset=0&limit=1", nil)
	router.ServeHTTP(w, findAllReq)
	albums := []model.Album{}
	json.NewDecoder(w.Body).Decode(&albums)

	// Fetch single album with previously retrieved id
	req, _ := http.NewRequest("GET", fmt.Sprintf("/album/%v", albums[0].ID), nil)
	router.ServeHTTP(w, req)

	album := model.Album{}

	if w.Code != 200 {
		t.Errorf("got %v, want %v", w.Code, 200)
	}

	json.NewDecoder(w.Body).Decode(&album)

	if len(album.Tracks) != 10 {
		t.Errorf("got %v, want %v", len(album.Tracks), 10)
	}
}

func TestSearch(t *testing.T) {
	router := router.NewRouter()

	// we assume that the following
	query := "accusantium"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/album?query=%s&offset=0&limit=10", query), nil)
	router.ServeHTTP(w, req)

	data := []model.Album{}
	json.NewDecoder(w.Body).Decode(&data)

	if len(data) != 10 {
		t.Errorf("got %v, want %v", len(data), 10)
	}
}
