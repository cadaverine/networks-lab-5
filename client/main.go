package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

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

	client := http.Client{}

	response, err := client.Get("http://" + ip)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
