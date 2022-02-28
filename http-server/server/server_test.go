package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var myprint = fmt.Println

// TestResponse tests basic response function
func TestResponse(t *testing.T) { // works
	// template declaration >> WebWriter(rw http.ResponseWriter, s string)
	req, err := http.NewRequest("GET", "/secret", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	hdl := http.HandlerFunc(Secret)

	hdl.ServeHTTP(res, req)

	// check status
	if status := res.Code; status != http.StatusOK {
		t.Errorf("status not ok bud :(")
	}

	// check body response
	expected := "psst, a secret\n"
	// fmt.Println("Response body: ", res.Body.String())
	if res.Body.String() != expected {
		t.Errorf("Body of response doesnt match really :(")

	}
}

// TestWebWriter tests webWriter function with & without parameter
// (does not use interfaces)
func TestWebWriter(t *testing.T) {
	res := httptest.NewRecorder()
	webWriter(res, "testing one")
	if res.Body.String() != "testing one\n" {
		t.Errorf("Wrong response, no parameter. Expected: %v; got: %v", "testing one\n", res.Body.String())
	}
	// fresh recorder
	res = httptest.NewRecorder()

	paramTest := "parameter value is me"
	SetParam(paramTest)
	expPreParam := "I have a parameter! Here: "

	writeTest := "testing 2"
	webWriter(res, writeTest)

	exp := (writeTest + "\n" + expPreParam + paramTest + "\n")
	if res.Body.String() != exp {
		t.Errorf("Wrong response, with parameter. Expected: %v; got: %v", exp, res.Body.String())
	}

	//TODO test error case next
}

// TestApiHome test api call to home page "/"
// It is expected that current time will be printed out
// func TestApiHome(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	hdlr := http.HandlerFunc(ApiHome)
// 	hdlr.ServeHTTP(rr, req)

// 	fmt.Println("status:", rr.Code)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	expected := time.Now().Format("2006-01-02 15:04:05")
// 	got := rr.Body.String()

// 	fmt.Println("exp:", expected, "; got:", got)
// 	if got != expected {

// 		t.Errorf("Http body fail. Expected %v but got %v ", expected, got)
// 	}
// }
