package artist

import (
	"allsounds/pkg/db"
	"allsounds/pkg/migration"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"allsounds/pkg/model"
)

var router *gin.Engine

func init() {
	db.DBCon, _ = gorm.Open(sqlite.Open(":memory:"), nil)
	migration.CreateTables()
	artists := migration.BulkInsertArtists(2)
	migration.BulkInsertAlbums(artists, 10)
	migration.BulkInsertUsers(10)

	router = gin.New()
	AddRoutes(router)
}

// TestFindAllArtistsWithoutPagination asserts that without offset or limit url parameters, endpoint will return 400
func TestFindAllArtistsWithoutPagination(t *testing.T) {
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

// TestFindAllArtists validates endpoint pagination
func TestFindAllArtists(t *testing.T) {
	w := httptest.NewRecorder()

	for i := 0; i <= 1; i += 1 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/artist?offset=%v&limit=1", i), nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		var data []model.Artist
		json.NewDecoder(w.Body).Decode(&data)

		if len(data) != 1 {
			t.Errorf("got %v, want %v", len(data), 1)
		}
	}
}

// TestAlbumById is a parameterized artist id test suite
func TestArtistById(t *testing.T) {
	var tests = []struct {
		name  string
		panic bool
	}{
		{name: "int"},
		{name: "string"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "int" {
				// First search in artist list to retrieve an actual artist id
				findAllReq, _ := http.NewRequest("GET", "/artist?offset=0&limit=1", nil)

				w := httptest.NewRecorder()

				router.ServeHTTP(w, findAllReq)
				var artists []model.Artist
				err := json.NewDecoder(w.Body).Decode(&artists)
				if err != nil {
					t.Errorf(err.Error())
				}

				// Fetch single artist with previously retrieved id
				req, _ := http.NewRequest("GET", fmt.Sprintf("/artist/%v", artists[0].ID), nil)
				router.ServeHTTP(w, req)

				artist := model.Artist{}

				if w.Code != 200 {
					t.Errorf("got %v, want %v", w.Code, 200)
				}

				err = json.NewDecoder(w.Body).Decode(&artist)
				if err != nil {
					t.Errorf(err.Error())
				}

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

// TestArtistSearch validates the search endpoint
func TestArtistSearch(t *testing.T) {
	query := ""
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/artist?query=%s&offset=0&limit=10", query), nil)
	router.ServeHTTP(w, req)

	var data []model.Artist
	err := json.NewDecoder(w.Body).Decode(&data)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(data) != 2 {
		t.Errorf("got %v, want %v", len(data), 2)
	}
}
