package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"allsounds/internal/dockertest"
	"allsounds/pkg/migration"

	"github.com/magiconair/properties/assert"
)

func TestAlbumsWithoutPagination(t *testing.T) {
	_, fncleanup := dockertest.SetupDockerDb()
	migration.CreateTables()

	router := NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/album", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
	assert.Equal(t, w.Body.String(), "{\"message\":\"bad request\"}")

	fncleanup()
}
