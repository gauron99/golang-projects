package main

import (
	"fmt"
	"math/rand"

	// "math/rand" // random number generator
	"net/http" //web stuff
	"net/url"

	"flag"
	"strconv" //string conversion
	"strings" //string manipulation
	"time"
)

var greetings = []string{"Hello", "Greetings", "Welcome", "Hi"}
var titles = []string{"the Mighty Traveler", "the Great Summoner", "the Conqueror of Titans", "the Destroyer of Worlds", "the King of Oceans", "the King of Underworld"}

var pageAccessCount = 0
var param string

// --- Parse cli arguments ---
// Only looking for long: "param" / short: "p"
// Can be given with "=" or as following argument
// 		ex: "--param=hello" or "--param hello" or "-p=hello"
// If param is given multiple times, save only the last one

func init() {

	flag.StringVar(&param, "param", "", "print this everywhere")
	flag.StringVar(&param, "p", "", "print this everywhere")

}

// write to browser (if ENV variable or CLI parameter are given, print those too)
func webWriter(rw http.ResponseWriter, s string) {
	fmt.Fprintf(rw, "%s\n", s)
	if len(param) > 0 { //if a parameter is given, print it out everywhere where something is printed out
		fmt.Fprintf(rw, "I have a parameter! Here: %s\n", param)
	}
}

// sayHello is a handler for api call "/hello"
// Its possible to give a 'name' parameter as "/hello?name=John"
// to greet John specificaly
func sayHello(rw http.ResponseWriter, req *http.Request) {

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
			if len(val) < 1 { //no name given
				webWriter(rw, "Greetings Mr. Nobody!")
			} else { // atleast one name given
				for _, name := range val {
					webWriter(rw, greetings[rand.Intn(len(greetings))]+" "+name+" "+titles[rand.Intn(len(titles))])
				}
			}
		}
	}
	if nameExists == false {
		webWriter(rw, "Hello there stranger! This page has been accessed "+strconv.FormatInt(int64(pageAccessCount), 10)+"x times\n")
	}
	pageAccessCount += 1
}

func apiHome(rw http.ResponseWriter, req *http.Request) {
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

func main() {

	// handle possible cli arguments & set
	flag.Parse()
	fmt.Println(param)

	fmt.Println("Server started...")
	defer fmt.Println("Server closed. Bye") //print at the end? need to catch signal to properly end main() probably

	http.HandleFunc("/", apiHome)
	http.HandleFunc("/hello", sayHello)

	http.ListenAndServe(":8000", nil)

}
