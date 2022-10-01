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
func TestAlbumsWithoutPagination(t *testing.T) {
	migration.CreateTables()

	router := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
	assert.Equal(t, w.Body.String(), "{\"message\":\"bad request\"}")
}

// Test endpoint pagination
func TestAlbums(t *testing.T) {
	migration.CreateTables()
	migration.BulkInsertAlbums()

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
