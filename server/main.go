package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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

	ip := host + ":" + port

	http.Handle("/", http.FileServer(http.Dir("static")))

	go serve(ip)

	fmt.Println("Server listening on " + ip)
	fmt.Scanln()
}
