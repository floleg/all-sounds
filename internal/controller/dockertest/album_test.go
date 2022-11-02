// Package dockertest implements end-to-end integration tests
// on the http routes layer
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

// TestFindAllAlbumsWithoutPagination asserts that without offset or limit url parameters, endpoint will return 400
func TestFindAllAlbumsWithoutPagination(t *testing.T) {
	testRouter := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album", nil)
	testRouter.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("got %v, want %v", w.Code, 400)
	}

	if w.Body.String() != "{\"message\":\"bad request\"}" {
		t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
	}
}

// TestFindAllAlbums validates endpoint pagination
func TestFindAllAlbums(t *testing.T) {
	testRouter := router.NewRouter()

	w := httptest.NewRecorder()

	for i := 0; i <= 9; i += 10 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/album?offset=%v&limit=10", i), nil)
		testRouter.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		var data []model.Album
		err := json.NewDecoder(w.Body).Decode(&data)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(data) != 10 {
			t.Errorf("got %v, want %v", len(data), 10)
		}
	}
}

// TestAlbumById is a parameterized album id test suite
func TestAlbumById(t *testing.T) {
	var tests = []struct {
		name  string
		panic bool
	}{
		{name: "int"},
		{name: "string"},
	}

	testRouter := router.NewRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "int" {
				// First search in album list to retrieve an actual album id
				findAllReq, _ := http.NewRequest("GET", "/album?offset=0&limit=1", nil)

				w := httptest.NewRecorder()

				testRouter.ServeHTTP(w, findAllReq)
				albums := []model.Album{}
				err := json.NewDecoder(w.Body).Decode(&albums)
				if err != nil {
					t.Errorf(err.Error())
				}

				// Fetch single album with previously retrieved id
				req, _ := http.NewRequest("GET", fmt.Sprintf("/album/%v", albums[0].ID), nil)
				testRouter.ServeHTTP(w, req)

				album := model.Album{}

				if w.Code != 200 {
					t.Errorf("got %v, want %v", w.Code, 200)
				}

				err = json.NewDecoder(w.Body).Decode(&album)
				if err != nil {
					t.Errorf(err.Error())
				}

				if len(album.Tracks) != 10 {
					t.Errorf("got %v, want %v", len(album.Tracks), 10)
				}
			} else {
				req, _ := http.NewRequest("GET", "/album/misguided-id", nil)

				w := httptest.NewRecorder()

				testRouter.ServeHTTP(w, req)

				if w.Code != 400 {
					t.Errorf("got %v, want %v", w.Code, 400)
				}

				if w.Body.String() != "{\"message\":\"bad request\"}" {
					t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
				}
			}
		})
	}
}

// TestAlbumsSearch validates the search endpoint
func TestAlbumsSearch(t *testing.T) {
	testRouter := router.NewRouter()

	query := "accusantium"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/album?query=%s&offset=0&limit=10", query), nil)
	testRouter.ServeHTTP(w, req)

	var data []model.Album
	err := json.NewDecoder(w.Body).Decode(&data)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(data) != 10 {
		t.Errorf("got %v, want %v", len(data), 10)
	}
}
