package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var greetings_array = []string{"Hello", "Greetings", "Welcome", "Hi"}
var titles_array = []string{"the Mighty Traveler", "the Great Summoner", "the Conqueror of Titans", "the Destroyer of Worlds", "the King of Oceans", "the King of Underworld"}

var page_access_count = 0

// sayHello is a handler for api call "/hello"
// Its possible to give a 'name' parameter as "/hello?name=John"
// to greet John specificaly
func sayHello(rw http.ResponseWriter, req *http.Request) {

	parser := (req.URL).String()
	if strings.Contains(parser, "name=") { //greet a person if "name" parameter is given
		name := parser[12:]
		if len(name) == 0 { //greet "nobody"
			fmt.Fprintf(rw, "Greetings Mr. Nobody!\n")
		} else { // greet specific person
			fmt.Fprintf(rw, "%s %s, %s!\n", greetings_array[rand.Intn(len(greetings_array))], name, titles_array[rand.Intn(len(titles_array))])
		}
	} else { //greet generic
		fmt.Fprintf(rw, "Hello there! This page has been accessed %dx times\n", page_access_count)
	}

	page_access_count += 1
}

func api_root(rw http.ResponseWriter, req *http.Request) {
	// start ticker
	// tick_sec := time.NewTicker(time.Second) //tick every second to show current time
	// done_sec := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-done_sec:
	// 			return
	// 		case t := <-tick_sec.C:
	// 			// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the pattern with which to format/parse a given time/string.
	// 			// display_time(t, rw)
	// 			fmt.Fprint(rw, "Now: ", t.Format("2006-01-02 15:04:05"))
	// 		}
	// 	}
	// }()

	t := time.Now()
	fmt.Fprint(rw, "Now: ", t.Format("2006-01-02 15:04:05"))
}

func main() {
	fmt.Println("Starting serer...")

	http.HandleFunc("/", api_root)
	http.HandleFunc("/hello", say_hello)

	http.ListenAndServe(":8000", nil)

}
