package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWebServerHandleTime(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/time", nil)
	s.handleTime().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), time.Now().Format(time.RFC1123))
}
