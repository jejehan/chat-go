package main

import (
	"log"
	"net/http"
)

func main() {

	var port = ":9090"

	//start webserver
	log.Println("ListenAndServe: localhost:", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("ListenAndServe:", err)
	}

}
