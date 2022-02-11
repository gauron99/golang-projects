package main

import (
	"fmt"
	"math/rand" // random number generator
	"net/http"  //web stuff
	"os"
	"strconv" //string conversion
	"strings" //string manipulation
	"time"
)

var greetings_array = []string{"Hello", "Greetings", "Welcome", "Hi"}
var titles_array = []string{"the Mighty Traveler", "the Great Summoner", "the Conqueror of Titans", "the Destroyer of Worlds", "the King of Oceans", "the King of Underworld"}

var page_access_count = 0
var param = ""

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

	parser := (req.URL).String()
	if strings.Contains(parser, "name=") { //greet a person if "name" parameter is given
		name := parser[12:]
		if len(name) == 0 { //greet "nobody"
			webWriter(rw, "Greetings Mr. Nobody!")
		} else { // greet specific person
			webWriter(rw, greetings_array[rand.Intn(len(greetings_array))]+" "+name+" "+titles_array[rand.Intn(len(titles_array))])
		}
	} else { //greet generic
		webWriter(rw, "Hello there! This page has been accessed "+strconv.FormatInt(int64(page_access_count), 10)+"x times\n")
	}

	page_access_count += 1

}

// TODO
// func getCurrTime() {
// }

func apiHome(rw http.ResponseWriter, req *http.Request) {
	// start ticker
	// tick_sec := time.NewTicker(1 * time.Second) //tick every second to show current time
	// for {
	// 	select {
	// 	case t := <-tick_sec.C:
	// 		fmt.Fprintf(rw, t.Format("2006-01-02 15:04:05"))
	// 	}
	// }
	t := time.Now()
	fmt.Println("Now: ", t.Format("2006-01-02 15:04:05"))
}

// --- Parse cli arguments ---
// Only looking for long: "parameter" or "param" / short: "p"
// Can be given with "=" or as following argument
// 		ex: "--param=hello" or "--param hello" or "-p=hello"
// Ignores unknown arguments
func argParser() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		capture_para := false

		for _, item := range args { //parse each argument
			if capture_para == true {
				param = item
				capture_para = false

			} else if strings.HasPrefix(item, "--parameter") || strings.HasPrefix(item, "--param") || strings.HasPrefix(item, "-p") {
				if strings.Contains(item, "=") { //if parameter contains '=' -> extract the argument
					item = item[strings.IndexByte(item, '='):][1:] //cut string until '=' (exluding =)

					if len(item) > 0 { //if len is 0 do nothing
						param = item
					}
				} else {
					capture_para = true // capture next parameter
				}

			} else {
				fmt.Println("Unknown parameter '", item, "' given, ignoring.")
			}
		}
	}
}

func main() {

	// handle possible cli arguments & set
	argParser()

	if param != "" {
		fmt.Println("I have a parameter!", param)
	}
	fmt.Println("Server started...")
	defer fmt.Println("Server closed. Bye") //print at the end? need to catch signal to properly end main() probably

	http.HandleFunc("/", apiHome)
	http.HandleFunc("/hello", sayHello)

	http.ListenAndServe(":8000", nil)

}
