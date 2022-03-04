package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
