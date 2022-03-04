package main

import (
	// "html/template" // for later use with html
	"flag"
	"fmt"
	"log"
	"net/http"

	"server/server"
)

func main() {
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
