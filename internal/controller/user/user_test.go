package user

import (
	"allsounds/internal/controller/track"
	"allsounds/pkg/migration"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"allsounds/pkg/db"
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
	track.AddRoutes(router)
}

// TestFindAllUsersWithoutPagination asserts that without offset or limit url parameters, endpoint will return 400
func TestFindAllUsersWithoutPagination(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("got %v, want %v", w.Code, 400)
	}

	if w.Body.String() != "{\"message\":\"bad request\"}" {
		t.Errorf("got %v, want %v", w.Body.String(), "{\"message\":\"bad request\"}")
	}
}

// TestFindAllUsers validates endpoint pagination
func TestFindAllUsers(t *testing.T) {
	w := httptest.NewRecorder()

	for i := 0; i <= 5; i += 1 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/user?offset=%v&limit=2", i), nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		data := []model.User{}
		json.NewDecoder(w.Body).Decode(&data)

		if len(data) != 2 {
			t.Errorf("got %v, want %v", len(data), 2)
		}
	}
}

// TestUserById is a parameterized artist id test suite
func TestUserById(t *testing.T) {
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
				// First search in user list to retreieve an actual user id
				findAllReq, _ := http.NewRequest("GET", "/user?offset=0&limit=1", nil)

				w := httptest.NewRecorder()

				router.ServeHTTP(w, findAllReq)
				users := []model.User{}
				json.NewDecoder(w.Body).Decode(&users)

				// Fetch single user with previously retrieved id
				req, _ := http.NewRequest("GET", fmt.Sprintf("/user/%v", users[0].ID), nil)
				router.ServeHTTP(w, req)

				if w.Code != 200 {
					t.Errorf("got %v, want %v", w.Code, 200)
				}
			} else {
				req, _ := http.NewRequest("GET", "/user/misguided-id", nil)

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

// TestUserSearch validates the search endpoint
func TestUserSearch(t *testing.T) {
	query := ""
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/user?query=%s&offset=0&limit=10", query), nil)
	router.ServeHTTP(w, req)

	var data []model.User
	err := json.NewDecoder(w.Body).Decode(&data)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(data) != 10 {
		t.Errorf("got %v, want %v", len(data), 10)
	}
}

func TestAppendUserTrack(t *testing.T) {
	// First search in user list to retrieve an actual user id
	findUsersReq, _ := http.NewRequest("GET", "/user?offset=0&limit=1", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, findUsersReq)
	users := []model.User{}
	err := json.NewDecoder(w.Body).Decode(&users)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Retrieve a list of tracks to append
	findTracksReq, _ := http.NewRequest("GET", "/track?offset=0&limit=10", nil)
	router.ServeHTTP(w, findTracksReq)
	var tracks []model.Track
	err = json.NewDecoder(w.Body).Decode(&tracks)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Append 10 tracks to user
	for _, track := range tracks {
		appendReq, _ := http.NewRequest("POST", fmt.Sprintf("/user/%v/track/%v", users[0].ID, track.ID), nil)
		router.ServeHTTP(w, appendReq)
	}

	// Assert tracks have been added to current user
	var user model.User
	_ = db.DBCon.Model(&model.User{}).Preload("Tracks").First(&user, users[0].ID).Error

	if len(user.Tracks) != 10 {
		t.Errorf("got %v, want %v", len(user.Tracks), 10)
	}
}
