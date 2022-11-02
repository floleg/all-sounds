package dockertest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"allsounds/internal/router"
	"allsounds/pkg/db"
	"allsounds/pkg/model"
)

// TestFindAllTracksWithoutPagination asserts that without offset or limit url parameters, endpoint will return 400
func TestFindAllTracksWithoutPagination(t *testing.T) {
	testRouter := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/track", nil)
	testRouter.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("got %v, want %v", w.Code, 400)
	}

	if w.Body.String() != "{\"message\":\"bad request\"}" {
		t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
	}
}

// TestFindAllTracks validates endpoint pagination
func TestFindAllTracks(t *testing.T) {
	testRouter := router.NewRouter()

	w := httptest.NewRecorder()

	for i := 0; i <= 100; i += 10 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/track?offset=%v&limit=10", i), nil)
		testRouter.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		var data []model.Track
		err := json.NewDecoder(w.Body).Decode(&data)

		if err != nil {
			t.Errorf(err.Error())
		}

		if len(data) != 10 {
			t.Errorf("got %v, want %v", len(data), 10)
		}
	}
}

// TestTrackById is a parameterized artist id test suite
func TestTrackById(t *testing.T) {
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
				// First search in track list to retrieve an actual track id
				findAllReq, _ := http.NewRequest("GET", "/track?offset=0&limit=1", nil)

				w := httptest.NewRecorder()

				testRouter.ServeHTTP(w, findAllReq)

				var tracks []model.Track
				err := json.NewDecoder(w.Body).Decode(&tracks)
				if err != nil {
					t.Errorf(err.Error())
				}

				// Fetch single track with previously retrieved id
				var track model.Track
				_ = db.DBCon.Model(&model.Track{}).Preload("Albums").First(&track, tracks[0].ID).Error

				if len(track.Albums) < 1 {
					t.Errorf("got %v, want an inflated albums slice", len(track.Albums))
				}
			} else {
				req, _ := http.NewRequest("GET", "/track/misguided-id", nil)

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

// TestTrackSearch validates the search endpoint
func TestTrackSearch(t *testing.T) {
	testRouter := router.NewRouter()

	query := ""
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/track?query=%s&offset=0&limit=10", query), nil)
	testRouter.ServeHTTP(w, req)

	var data []model.Track
	err := json.NewDecoder(w.Body).Decode(&data)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(data) != 10 {
		t.Errorf("got %v, want %v", len(data), 10)
	}
}
