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

// Without offset or limit url parameters, endpoint will return 400
func TestFindAllTracksWithoutPagination(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/track", nil)
	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("got %v, want %v", w.Code, 400)
	}

	if w.Body.String() != "{\"message\":\"bad request\"}" {
		t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
	}
}

// Test endpoint pagination
func TestFindAllTracks(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()

	for i := 0; i <= 100; i += 10 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/track?offset=%v&limit=10", i), nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		data := []model.Track{}
		json.NewDecoder(w.Body).Decode(&data)

		if len(data) != 10 {
			t.Errorf("got %v, want %v", len(data), 10)
		}
	}
}

// Two cases: integer parameter and string parameter
func TestTrackById(t *testing.T) {
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
				// First search in track list to retreieve an actual track id
				findAllReq, _ := http.NewRequest("GET", "/track?offset=0&limit=1", nil)

				w := httptest.NewRecorder()

				router.ServeHTTP(w, findAllReq)
				tracks := []model.Track{}
				json.NewDecoder(w.Body).Decode(&tracks)

				// Fetch single track with previously retrieved id
				var track model.Track
				_ = db.DBCon.Model(&model.Track{}).Preload("Albums").First(&track, tracks[0].ID).Error

				if len(track.Albums) < 1 {
					t.Errorf("got %v, want an inflated albums slice", len(track.Albums))
				}
			} else {
				req, _ := http.NewRequest("GET", "/track/misguided-id", nil)

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

func TestSearchTrack(t *testing.T) {
	router := router.NewRouter()

	query := ""
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/track?query=%s&offset=0&limit=10", query), nil)
	router.ServeHTTP(w, req)

	data := []model.Track{}
	json.NewDecoder(w.Body).Decode(&data)

	if len(data) != 10 {
		t.Errorf("got %v, want %v", len(data), 10)
	}
}
