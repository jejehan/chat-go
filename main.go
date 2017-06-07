package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("ListenAndServe: localhost:9090")

	//start webserver
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println("ListenAndServe:", err)
	}

}
