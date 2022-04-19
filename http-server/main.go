package main

import (
	// "html/template" // for later use with html

	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"server/server"
)

func main() {
	// default PORT value
	var PORT = "8000"
	var paramPtr string
	var portPtr string
	flag.StringVar(&paramPtr, "param", "", "print this everywhere(long)")
	flag.StringVar(&paramPtr, "p", "", "print this everywhere(short)")
	flag.StringVar(&portPtr, "port", "", "add port to run on (default="+PORT+")")

	flag.Parse()

	// have default parameter set as ENV variable if --param is not given
	if paramPtr == "" {
		var check bool
		paramPtr, check = os.LookupEnv("SERVER_PARAM")
		if check {
			// just use check so it doesnt return an error :grin:
		}
	}

	// check if port was given
	if portPtr != "" {
		PORT = portPtr
	}

	serv, e := server.NewServerInfo(paramPtr)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Server started...")

	http.HandleFunc("/", serv.ApiHome)
	http.HandleFunc("/hello", serv.SayHello)

	log.Println("Listening on port: ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}
