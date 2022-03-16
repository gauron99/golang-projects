package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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

	s, e := NewServerInfo("")
	if e != nil {
		log.Fatal(e)
	}

	paramPre := "I have a parameter! Here: "
	envVarPre := "\nMy environment variables: "

	for _, data := range testsWebWriterMetaData {
		t.Run(data.name, func(t *testing.T) {
			res := httptest.NewRecorder()
			// set parameter
			s.param = data.metadata["param"].(string)
			s.variables = data.metadata["envVar"].(map[string]string)

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

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////// TESTING APIHOME ///////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// TestApiHome test ApiHome handler for "/"
func TestApiHome(t *testing.T) {

	s, e := NewServerInfo("")
	if e != nil {
		log.Fatal(e)
	}

	envVarPre := "\nMy environment variables: "

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	s.ApiHome(res, req)
	exp := time.Now().Format("2006-01-02 15:04:05") + "\n"

	if len(s.variables) > 0 {
		exp += fmt.Sprintf("%s%v\n", envVarPre, s.variables)
	}

	got := res.Body.String()
	if got != exp {
		t.Errorf("got |%s|, but expected |%s|", got, exp)
	}
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////// TESTING WEBWRITER //////////////////////////////
///////////////////////////////////////////////////////////////////////////////

type testsSayHelloMetadataStruct struct {
	name   string
	str    string
	helper string
}

var testsSayHelloData = []testsSayHelloMetadataStruct{
	{
		name:   "nothing given, return stranger",
		str:    "",
		helper: "stranger",
	},
	{
		name:   "just empty name",
		str:    "?name",
		helper: "nobody",
	},
	{
		name:   "just empty name (with equal sign)",
		str:    "?name=",
		helper: "nobody",
	},
	{
		name:   "one name given",
		str:    "?name=Jonathan",
		helper: "Jonathan",
	},
	{
		name:   "2 names given",
		str:    "?name=David&name=Goliath",
		helper: "David,Goliath",
	},
}

// TestSayHello tests
func TestSayHello(t *testing.T) {
	s, e := NewServerInfo("")
	if e != nil {
		log.Fatal(e)
	}

	envVarPre := "\nMy environment variables: "

	// accessCnt := 0
	// stranger := fmt.Sprintf("Hello there stranger! This page has been accessed " + strconv.FormatInt(int64(accessCnt), 10) + "x times")
	stranger := "Hello there stranger! This page has been accessed 0x times"
	nobody := "Greetings Mr. Nobody!"

	for _, data := range testsSayHelloData {
		t.Run(data.name, func(t *testing.T) {
			urlIn := "/hello" + data.str
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, urlIn, nil)

			s.SayHello(res, req)
			got := res.Body.String()
			exp := ""

			sentence := strings.Split(got, "\n")[0]
			sentences := strings.Split(sentence, " and ")
			sentenceConnector := " and "

			// setting expected value
			if data.helper == "nobody" {
				exp = nobody
			} else if data.helper == "stranger" {
				exp = stranger
			} else { //name given! -- check immediately
				var isOK bool
				for i, name := range strings.Split(data.helper, ",") {
					isOK = false
					for _, greeting := range greetings { //check greetings
						if strings.HasPrefix(sentences[i], greeting+" "+name) {
							for _, title := range titles { // check title
								if sentences[i] == greeting+" "+name+" "+title {
									isOK = true
									exp += sentences[i]
									if i < len(sentences)-1 { //connect names in sentence
										exp += sentenceConnector
									}
									break
								}
							}
						}
					}
				}
				if !isOK {
					if exp == "" {
						t.Errorf("Strings did not match with names given, got: %s", got)
					} else {
						t.Errorf("Got: %s, but expected: %s.", got, exp)
					}
				}

			} //end of else
			exp = exp + "\n"
			if len(s.variables) > 0 {
				exp += fmt.Sprintf("%s%v\n", envVarPre, s.variables)
			}

			if exp != got {
				t.Errorf("got:\n%s; exp:\n%s", got, exp)
			}

		})

	}
}
