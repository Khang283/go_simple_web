package test

import (
	"crud/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(router.HandlerHello)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler return wrong status code")
	}
}
