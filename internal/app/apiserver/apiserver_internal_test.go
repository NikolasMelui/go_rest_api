package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleMain(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	s.handleMain().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Main handler")
}

func TestAPIServer_HandleFirst(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/first", nil)
	s.handleFirst().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "First handler")
}

func TestAPIServer_HandleSecond(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/second", nil)
	s.handleSecond().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Second handler")
}
