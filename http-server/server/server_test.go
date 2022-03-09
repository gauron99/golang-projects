package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
////////////////////////////// TESTING WEBWRITER //////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type webWriterMetadataTestStruct struct {
	name     string
	metadata map[string]string //input, parameter, environment var
	isError  bool
}

var testsWebWriterMetaData = []webWriterMetadataTestStruct{
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
			"input":  "input written",
			"param":  "",
			"envVar": "",
		},
		isError: false,
	},
	{
		name: "input + param given, should pass",
		metadata: map[string]string{
			"input":  "input written",
			"param":  "param given",
			"envVar": "",
		},
		isError: false,
	},
}

// TestWebWriter tests webWriter function with & without parameter
// (does not use interfaces)
func TestWebWriter(t *testing.T) {
	os.Clearenv()
	s := NewServerInfo("")

	paramPre := "I have a parameter! Here: "
	envVarPre := "\nMy environment variables: "

	for _, data := range testsWebWriterMetaData {
		t.Run(data.name, func(t *testing.T) {
			res := httptest.NewRecorder()
			s.param = data.metadata["param"]

			s.webWriter(res, data.metadata["input"])
			if res.Result().StatusCode != http.StatusOK {
				t.Error("Expected status '200 OK' but got:", res.Result().Status)
			}

			exp := ""
			if data.metadata["input"] != "" {
				exp += data.metadata["input"] + "\n"
			}
			if data.metadata["param"] != "" {
				exp += paramPre + data.metadata["param"] + "\n"
			}
			if data.metadata["envVar"] != "" {
				exp += envVarPre + data.metadata["envVar"] + "\n"
			}

			if res.Body.String() != exp {
				t.Error("data printed out (", res.Body.String(), ") does not match the expected (", exp, ")")
			}

		})
	}
}

// TestApiHome test ApiHome handler for "/"
func TestApiHome(t *testing.T) {
	os.Clearenv()
	s := NewServerInfo("")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	s.ApiHome(res, req)
	exp := time.Now().Format("2006-01-02 15:04:05") + "\n"
	got := res.Body.String()
	if got != exp {
		t.Errorf("got |%s|, but expected |%s|", got, exp)
	}
}
