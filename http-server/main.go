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
	var paramPtr string
	flag.StringVar(&paramPtr, "param", "", "print this everywhere(long)")
	flag.StringVar(&paramPtr, "p", "", "print this everywhere(short)")

	flag.Parse()

	// have default parameter set as ENV variable if --param is not given
	if paramPtr == "" {
		var check bool
		paramPtr, check = os.LookupEnv("SERVER_PARAM")
		if check {
			// just use check so it doesnt return an error :grin:
		}
	}

	serv, e := server.NewServerInfo(paramPtr)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Server started...")

	http.HandleFunc("/", serv.ApiHome)
	http.HandleFunc("/hello", serv.SayHello)

	log.Fatal(http.ListenAndServe(":8001", nil))


}
