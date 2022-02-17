package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var pageAccessCount = 0
var param string

var greetings = []string{"Hello", "Greetings", "Welcome", "Hi"}
var titles = []string{"the Mighty Traveler", "the Great Summoner", "the Conqueror of Titans", "the Destroyer of Worlds", "the King of Oceans", "the King of Underworld"}

// GetParam returns pointer to param variable (for other packages)
func GetParam() *string {
	return &param
}

// webWriter takes input string and writes it to browser
// additionally if ENV variable or CLI parameter are given, print those too
func webWriter(rw http.ResponseWriter, s string) {
	fmt.Fprintf(rw, "%s\n", s)
	if len(param) > 0 { //if a parameter is given, print it out everywhere where something is printed out
		fmt.Fprintf(rw, "I have a parameter! Here: %s\n", param)
	}
}

// SayHello is a handler for api call "/hello"
// Its possible to give a 'name' parameter as "/hello?name=John"
// to greet John specificaly
func SayHello(rw http.ResponseWriter, req *http.Request) {

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
				webWriter(rw, "Greetings Mr. Nobody!")
			} else { // atleast one name given
				for _, name := range val {
					webWriter(rw, greetings[rand.Intn(len(greetings))]+" "+name+" "+titles[rand.Intn(len(titles))])
				}
			}
		}
	}
	if nameExists == false {
		webWriter(rw, "Hello there stranger! This page has been accessed "+strconv.FormatInt(int64(pageAccessCount), 10)+"x times")
	}
	pageAccessCount += 1
}

func ApiHome(rw http.ResponseWriter, req *http.Request) {
	// start ticker
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
