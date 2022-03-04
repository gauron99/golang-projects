package server

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// change this to use methods in the file
type serverInfo struct {
	pageAccessCount int
	Param           string
}

// MakeServerInfo creates & returns newly initialized serverInfo structure with given args
func MakeServerInfo(pageAccessCnt int, parameter string) *serverInfo {
	return &serverInfo{pageAccessCnt, parameter}
}

var greetings = []string{"Hello", "Greetings", "Welcome", "Hi"}
var titles = []string{"the Mighty Traveler", "the Great Summoner", "the Conqueror of Titans", "the Destroyer of Worlds", "the King of Oceans", "the King of Underworld"}

// webWriter takes input string and writes it to browser
// additionally if ENV variable or CLI parameter are given, print those too
func (si serverInfo) webWriter(rw http.ResponseWriter, s string) {
	fmt.Fprintf(rw, "%s\n", s)
	if len(si.Param) > 0 { //if a parameter is given, print it out everywhere where something is printed out

		out := fmt.Sprintf("I have a parameter! Here: %s\n", si.Param)
		io.WriteString(rw, out)
	}
}

// SayHello is a handler for api call "/hello"
// Its possible to give a 'name' parameter as "/hello?name=John"
// to greet John specificaly
func (si *serverInfo) SayHello(rw http.ResponseWriter, req *http.Request) {

	URL := req.URL.String()
	u, err := url.Parse(URL) //parsed url
	if err != nil {
		panic(err)
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
	t := time.Now()
	fmt.Println("Now: ", t.Format("2006-01-02 15:04:05"))
}
