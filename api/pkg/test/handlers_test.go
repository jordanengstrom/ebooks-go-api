package tests

import (
	"ebooks/pkg/db"
	"ebooks/pkg/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckReturnsPong(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/ping", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	DB := db.Init()
	h := handlers.New(DB)

	testHandler := http.HandlerFunc(h.HealthCheck)
	testHandler.ServeHTTP(rr, req)

	const expected string = "\"pong\"\n"
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body. Expected: %v Actual: %v", actual, expected)
	}
}

func TestHealthCheckReturns200(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/ping", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	DB := db.Init()
	h := handlers.New(DB)

	testHandler := http.HandlerFunc(h.HealthCheck)
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code. Expected: %v Actual: %v", status, http.StatusOK)
	}
}
