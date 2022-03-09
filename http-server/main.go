package main

import (
	// "html/template" // for later use with html
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"server/server"
)

// checkErr is a helper function for checking if return values are nil
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

// loadEnvSettings reads data from a file called "env" in current dir
// and sets all environment variables for this run.
// Returns slice of keys in the env file.
func loadEnvSettings() {

	//clear all pre-set variables since we dont need any
	os.Clearenv()
	f, err := os.Open("./env")
	checkErr(err)
	defer f.Close() // close when func is done

	scanner := bufio.NewScanner(f)

	// cycle through each line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue //lines starting with '#' are comments
		}
		if !strings.Contains(line, "=") { //line has to contain '='
			panic("line in env file doesn't contain '='")
		}
		split := strings.Split(line, "=")
		os.Setenv(split[0], split[1])
	}
}

func main() {
	loadEnvSettings()

	var paramPtr string
	flag.StringVar(&paramPtr, "param", "", "print this everywhere(long)")
	flag.StringVar(&paramPtr, "p", "", "print this everywhere(short)")
	// handle possible cli arguments & set
	flag.Parse()

	serv := server.NewServerInfo(paramPtr)

	fmt.Println("Server started...")

	http.HandleFunc("/", serv.ApiHome)
	http.HandleFunc("/hello", serv.SayHello)

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
