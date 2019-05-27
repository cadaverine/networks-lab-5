package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server response")
}

func serve(ip string) {
	log.Fatal(http.ListenAndServe(ip, nil))
}

func main() {
	host, port := "localhost", "8080"

	switch len(os.Args) {
	case 3:
		port = os.Args[2]
		host = os.Args[1]
	case 2:
		host = os.Args[1]
	}

	ipAddress := host + ":" + port

	http.HandleFunc("/", handler)

	go serve(ipAddress)

	fmt.Println("Server listening on " + ipAddress)
	fmt.Scanln()
}
