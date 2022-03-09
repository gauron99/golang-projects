package server

import (
	"fmt"
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
	metadata map[string]interface{} //input, parameter, environment var(map)
	isError  bool
}

var testsWebWriterMetaData = []webWriterMetadataTestStruct{
	{
		name: "nothing passed, should pass",
		metadata: map[string]interface{}{
			"input": "",
			"param": "",
			"envVar": map[string]string{
				"VAR": ""},
		},
		isError: false,
	},
	{
		name: "just input given, should pass",
		metadata: map[string]interface{}{
			"input": "input written",
			"param": "",
			"envVar": map[string]string{
				"VAR": ""},
		},
		isError: false,
	},
	{
		name: "param given, should pass",
		metadata: map[string]interface{}{
			"input": "",
			"param": "param given",
			"envVar": map[string]string{
				"VAR": ""},
		},
		isError: false,
	},
	{
		name: "input + param given, should pass",
		metadata: map[string]interface{}{
			"input": "input written",
			"param": "param given",
			"envVar": map[string]string{
				"VAR": ""},
		},
		isError: false,
	},
	{
		name: "envVar given, should pass",
		metadata: map[string]interface{}{
			"input": "",
			"param": "",
			"envVar": map[string]string{
				"VAR": "envVar given"},
		},
		isError: false,
	},
	{
		name: "input + envVar given, should pass",
		metadata: map[string]interface{}{
			"input": "input written",
			"param": "",
			"envVar": map[string]string{
				"VAR": "envVar given"},
		},
		isError: false,
	},
	{
		name: "param + envVar given, should pass",
		metadata: map[string]interface{}{
			"input": "",
			"param": "param given",
			"envVar": map[string]string{
				"VAR": "envVar given"},
		},
		isError: false,
	},
	{
		name: "input + param + envVar given, should pass",
		metadata: map[string]interface{}{
			"input": "input written",
			"param": "param given",
			"envVar": map[string]string{
				"VAR": "envVar given"},
		},
		isError: false,
	},
	{
		name: "input + param + envVar (multiple) given, should pass",
		metadata: map[string]interface{}{
			"input": "input written",
			"param": "param given",
			"envVar": map[string]string{
				"VAR": "hello",
				"BAR": "2",
				"FOO": "interface not really"},
		},
		isError: false,
	},
}

// TestWebWriter tests webWriter function with & without parameter
func TestWebWriter(t *testing.T) {
	s := NewServerInfo("")

	paramPre := "I have a parameter! Here: "
	envVarPre := "\nMy environment variables: "

	for _, data := range testsWebWriterMetaData {
		t.Run(data.name, func(t *testing.T) {
			res := httptest.NewRecorder()
			// set parameter
			s.param = data.metadata["param"].(string)

			// set environment vars
			os.Clearenv()
			for k, v := range data.metadata["envVar"].(map[string]string) {
				os.Setenv(k, v)
			}

			s.webWriter(res, data.metadata["input"].(string))
			if res.Result().StatusCode != http.StatusOK {
				t.Error("Expected status '200 OK' but got:", res.Result().Status)
			}

			exp := ""
			if data.metadata["input"] != "" {
				exp += data.metadata["input"].(string) + "\n"
			}
			if data.metadata["param"] != "" {
				exp += paramPre + data.metadata["param"].(string) + "\n"
			}
			if len(data.metadata["envVar"].(map[string]string)) > 0 {
				exp += fmt.Sprintf("%s%v\n", envVarPre, data.metadata["envVar"].(map[string]string))
			}

			if res.Body.String() != exp {
				t.Errorf("got |%s|, expected |%s|", res.Body.String(), exp)
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
