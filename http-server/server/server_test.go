package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
// TestWebWriter tests webWriter() function and its different input
// possibilities (with if statement for cli parameter or env variable)
func TestWebWriter(t *testing.T) {
	// template declaration >> WebWriter(rw http.ResponseWriter, s string)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	hdl := http.HandlerFunc()
}
*/

// TestApiHome test api call to home page "/"
// It is expected that current time will be printed out
func TestApiHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	hdlr := http.HandlerFunc(ApiHome)
	hdlr.ServeHTTP(rr, req)

	fmt.Println("status:", rr.Code)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := time.Now().Format("2006-01-02 15:04:05")
	got := rr.Body.String()

	fmt.Println("exp:", expected, "; got:", got)
	if got != expected {

		t.Errorf("Http body fail. Expected %v but got %v ", expected, got)
	}
}
