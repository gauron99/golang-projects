package main

import (
	// "html/template" // for later use with html
	"flag"
	"fmt"
	"log"
	"net/http"

	"server/server"
)

// --- Parse cli arguments ---
// Only looking for long: "param" / short: "p"
// Can be given with "=" or as following argument
// 		ex: "--param=hello" or "--param hello" or "-p=hello"
// If param is given multiple times, save only the last one

func init() {

	flag.StringVar(&server.Param, "param", "", "print this everywhere")
	flag.StringVar(&server.Param, "p", "", "print this everywhere")

}

func main() {

	// handle possible cli arguments & set
	flag.Parse()
	fmt.Println(server.Param)

	fmt.Println("Server started...")
	defer fmt.Println("Server closed. Bye") //print at the end? need to catch signal to properly end main() probably

	http.HandleFunc("/", server.ApiHome)
	http.HandleFunc("/hello", server.SayHello)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
