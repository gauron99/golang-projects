package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// change this to use methods in the file
type serverInfo struct {
	pageAccessCount int
	param           string
	variables       map[string]string
}

// LoadSettings reads data from a file called "env" in current dir
// and sets all environment variables for this run.
// Returns map of keys&values in the env file.
func loadSettings() (vars map[string]string, err error) {
	dir, err := os.Getwd()
	vars = make(map[string]string)
	var filName string
	if strings.HasSuffix(dir, "/server") {
		// trying to run this in server/server_test.go
		filName = "../env"
	} else {
		filName = "./env"
	}
	f, err := os.Open(filName)

	// prematurely return an error
	if err != nil {
		return nil, err
	}

	defer f.Close() // close when func is done

	scanner := bufio.NewScanner(f)

	// cycle through each line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue //lines starting with '#' are comments
		}
		if !strings.Contains(line, "=") { //line has to contain '='
			log.Println("Warning: env_file wasn't properly loaded; line in env file doesn't contain '=':", line)
			return vars, nil
		}
		split := strings.Split(line, "=")
		vars[split[0]] = split[1]
	}
	return vars, nil
}

// getEnvVars reads all env variables and returns them as a map (foo[key] = val)
func (si serverInfo) getEnvVars() map[string]string {
	res := make(map[string]string)
	for key, val := range si.variables {
		res[key] = val
	}
	return res
}

// MakeServerInfo creates & returns newly initialized serverInfo structure with given args
func NewServerInfo(parameter string) (*serverInfo, error) {
	vars, e := loadSettings()
	return &serverInfo{0, parameter, vars}, e
}

var greetings = []string{"Hello", "Greetings", "Welcome", "Hi"}
var titles = []string{"the Mighty Traveler", "the Great Summoner", "the Conqueror of Titans", "the Destroyer of Worlds", "the King of Oceans", "the King of Underworld"}

// webWriter takes input string and writes it to browser
// additionally if ENV variable or CLI parameter are given, print those too
func (si serverInfo) webWriter(rw http.ResponseWriter, s string) {
	if s != "" {
		outS := fmt.Sprintf("%s\n", s)
		_, err := io.WriteString(rw, outS)
		if err != nil {
			log.Printf("Error while writing string: %s", err)
		}
	}

	if par := si.param; len(par) > 0 { //if a parameter is given, print it out everywhere where something is printed out

		outParam := fmt.Sprintf("I have a parameter! Here: %s\n", par)
		_, err := io.WriteString(rw, outParam)
		if err != nil {
			log.Printf("Error while writing paramaters: %s", err)
		}
	}
	if vars := si.getEnvVars(); len(vars) > 0 {
		outVars := fmt.Sprintf("\nMy environment variables: %v\n", vars)
		_, err := io.WriteString(rw, outVars)
		if err != nil {
			log.Printf("Error while writing environment variables: %s", err)
		}
	}
}

// SayHello is a handler for api call "/hello"
// Its possible to give a 'name' parameter as "/hello?name=John"
// to greet John specificaly
func (si *serverInfo) SayHello(rw http.ResponseWriter, req *http.Request) {

	URL := req.URL.String()
	u, err := url.Parse(URL) //parsed url
	if err != nil {
		log.Printf("Error while parsing URL: %s", err)
		rw.WriteHeader(500)
	}
	// parse query
	args, _ := url.ParseQuery(u.RawQuery)

	nameExists := false
	// cycle through all arguments given and search for "name"
	for key, val := range args {
		if strings.Compare(key, "name") == 0 { //name argument exists
			nameExists = true
			if len(val) == 1 && len(val[0]) == 0 { //no name given
				si.webWriter(rw, "Greetings Mr. Nobody!")
			} else { // atleast one name given
				for _, name := range val {
					si.webWriter(rw, greetings[rand.Intn(len(greetings))]+" "+name+" "+titles[rand.Intn(len(titles))])
				}
			}
		}
	}
	if !nameExists {
		si.webWriter(rw, "Hello there stranger! This page has been accessed "+strconv.FormatInt(int64(si.pageAccessCount), 10)+"x times")
	}
	si.pageAccessCount += 1
}

func (si serverInfo) ApiHome(rw http.ResponseWriter, req *http.Request) {
	// TODO start ticker
	// tickSec := time.NewTicker(1 * time.Second) //tick every second to show current time
	// for {
	// 	select {
	// 	case t := <-tickSec.C:
	// 		fmt.Fprintf(rw, t.Format("2006-01-02 15:04:05"))
	// 	}
	// }
	out := time.Now().Format("2006-01-02 15:04:05")
	// out = strings.TrimSuffix(out, "\n")
	// fmt.Println("LALA", out)
	si.webWriter(rw, out)
}
