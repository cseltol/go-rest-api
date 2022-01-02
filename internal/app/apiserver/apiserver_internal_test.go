package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHi(t *testing.T) {
	s := New(NewConfig())
	rec :=  httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hi", nil)
	s.handleHi().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hi there!")
}
