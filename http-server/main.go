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

// loadEnvSettings reads data from a file called "env" in current dir
// and sets all environment variables for this run.
// Returns slice of keys in the env file.
func loadEnvSettings() (err error) {

	//clear all pre-set variables since we dont need any
	os.Clearenv()
	f, err := os.Open("./env")
	if err != nil {
		return err
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
			return nil
		}
		split := strings.Split(line, "=")
		os.Setenv(split[0], split[1])
	}
	return nil
}

func main() {
	e := loadEnvSettings()
	if e != nil {
		log.Fatal(e) //print out error and exit
	}

	os.Setenv("HARDCODED", "nejvic") //custom env var

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
