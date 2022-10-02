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
func TestFindAllAlbumsWithoutPagination(t *testing.T) {
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
func TestFindAllAlbums(t *testing.T) {
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

// Two cases: integer parameter and string parameter
func TestAlbumById(t *testing.T) {
	var tests = []struct {
		name  string
		panic bool
	}{
		{name: "int"},
		{name: "string"},
	}

	router := router.NewRouter()

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			if tt.name == "int" {
				// First search in album list to retreieve an actual album id
				findAllReq, _ := http.NewRequest("GET", "/album?offset=0&limit=1", nil)

				w := httptest.NewRecorder()

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
			} else {
				req, _ := http.NewRequest("GET", "/album/misguided-id", nil)

				w := httptest.NewRecorder()

				router.ServeHTTP(w, req)

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

func TestAlbumsSearch(t *testing.T) {
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
