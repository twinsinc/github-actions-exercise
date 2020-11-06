package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(root)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code returned: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Hello World.`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Unexpected body returned: got %v want %v", actual, expected)
	}
}
