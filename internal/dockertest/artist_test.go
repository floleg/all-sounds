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
func TestFindAllArtistsWithoutPagination(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/artist", nil)
	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("got %v, want %v", w.Code, 400)
	}

	if w.Body.String() != "{\"message\":\"bad request\"}" {
		t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
	}
}

// Test endpoint pagination
func TestFindAllArtists(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()

	for i := 0; i <= 1; i += 1 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/artist?offset=%v&limit=10", i), nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		data := []model.Artist{}
		json.NewDecoder(w.Body).Decode(&data)

		if len(data) != 2 {
			t.Errorf("got %v, want %v", len(data), 2)
		}
	}
}

// Two cases: integer parameter and string parameter
func TestArtistById(t *testing.T) {
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
				// First search in artist list to retreieve an actual artist id
				findAllReq, _ := http.NewRequest("GET", "/artist?offset=0&limit=1", nil)

				w := httptest.NewRecorder()

				router.ServeHTTP(w, findAllReq)
				artists := []model.Artist{}
				json.NewDecoder(w.Body).Decode(&artists)

				// Fetch single artist with previously retrieved id
				req, _ := http.NewRequest("GET", fmt.Sprintf("/artist/%v", artists[0].ID), nil)
				router.ServeHTTP(w, req)

				artist := model.Artist{}

				if w.Code != 200 {
					t.Errorf("got %v, want %v", w.Code, 200)
				}

				json.NewDecoder(w.Body).Decode(&artist)

				if len(artist.Tracks) == 0 {
					t.Errorf("got %v, want an inflated slice", len(artist.Tracks))
				}
			} else {
				req, _ := http.NewRequest("GET", "/artist/misguided-id", nil)

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

func TestSearchArtist(t *testing.T) {
	router := router.NewRouter()

	query := ""
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/artist?query=%s&offset=0&limit=10", query), nil)
	router.ServeHTTP(w, req)

	data := []model.Artist{}
	json.NewDecoder(w.Body).Decode(&data)

	if len(data) != 2 {
		t.Errorf("got %v, want %v", len(data), 2)
	}
}
