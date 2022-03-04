package server

import (
	"net/http/httptest"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////
////////////////////////////// TESTING WEBWRITER //////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type webWriterMetadataTestStruct struct {
	name     string
	metadata map[string]string //input, parameter, environment var
	isError  bool
}

var testsWebWriterMeta = []webWriterMetadataTestStruct{
	{
		name: "nothing passed, should pass",
		metadata: map[string]string{
			"input":  "",
			"param":  "",
			"envVar": "",
		},
		isError: false,
	},
	{
		name: "just input given, should pass",
		metadata: map[string]string{
			"input":  "write this",
			"param":  "",
			"envVar": "",
		},
		isError: false,
	},
	{
		name: "input + param given, should pass",
		metadata: map[string]string{
			"input":  "",
			"param":  "",
			"envVar": "",
		},
		isError: false,
	},
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
	Serv.SetParam(paramTest)
	expPreParam := "I have a parameter! Here: "

	writeTest := "testing 2"
	webWriter(res, writeTest)

	exp := (writeTest + "\n" + expPreParam + paramTest + "\n")
	if res.Body.String() != exp {
		t.Errorf("Wrong response, with parameter. Expected: %v; got: %v", exp, res.Body.String())
	}

	//TODO test error case next
}
