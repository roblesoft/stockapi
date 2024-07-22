package controller_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/roblesoft/stockapi/internal/controller/http"
)

func TestGetPing_Success(t *testing.T) {
	server := controller.NewServer()
	server.SetupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	server.Ping(w, req)

	res := w.Result()

	defer res.Body.Close()

	_, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected to be nil got %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status ok got %d", res.StatusCode)
	}
}
