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
func TestFindAllUsersWithoutPagination(t *testing.T) {
	router := router.NewRouter()

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

// Test endpoint pagination
func TestFindAllUsers(t *testing.T) {
	router := router.NewRouter()

	w := httptest.NewRecorder()

	for i := 0; i <= 100; i += 10 {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/user?offset=%v&limit=10", i), nil)
		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("got %v, want %v", w.Code, 200)
		}

		data := []model.User{}
		json.NewDecoder(w.Body).Decode(&data)

		if len(data) != 10 {
			t.Errorf("got %v, want %v", len(data), 10)
		}
	}
}

// Two cases: integer parameter and string parameter
func TestUserById(t *testing.T) {
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

func TestSearchUser(t *testing.T) {
	router := router.NewRouter()

	query := ""
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/user?query=%s&offset=0&limit=10", query), nil)
	router.ServeHTTP(w, req)

	data := []model.User{}
	json.NewDecoder(w.Body).Decode(&data)

	if len(data) != 10 {
		t.Errorf("got %v, want %v", len(data), 10)
	}
}

func TestAppendUserTrack(t *testing.T) {
	router := router.NewRouter()

	// First search in user list to retreieve an actual user id
	findUsersReq, _ := http.NewRequest("GET", "/user?offset=0&limit=1", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, findUsersReq)
	users := []model.User{}
	json.NewDecoder(w.Body).Decode(&users)

	// Retrieve a list of tracks to append
	findTracksReq, _ := http.NewRequest("GET", "/track?offset=0&limit=10", nil)
	router.ServeHTTP(w, findTracksReq)
	tracks := []model.Track{}
	json.NewDecoder(w.Body).Decode(&tracks)

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
