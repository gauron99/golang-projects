package main

import (
	// "html/template" // for later use with html
	"flag"
	"fmt"
	"log"
	"net/http"

	"server/server"
)

var paramPtr string

// --- Parse cli arguments ---
// Only looking for long: "param" / short: "p"
// Can be given with "=" or as a following argument
// 		ex: "--param=hello" or "--param hello" or "-p=hello"
// If param is given multiple times, save only the last one

func init() {
	flag.StringVar(&paramPtr, "param", "", "print this everywhere(long)")
	flag.StringVar(&paramPtr, "p", "", "print this everywhere(short)")
}

func main() {

	// handle possible cli arguments & set
	flag.Parse()
	server.Serv.SetParam(paramPtr)

	fmt.Println("Server started...")

	http.HandleFunc("/", server.ApiHome)
	http.HandleFunc("/hello", server.SayHello)
	http.HandleFunc("/secret", server.Secret)

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
