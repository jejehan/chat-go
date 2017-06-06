package main

import (
	"log"
	"net/http"
)

func main() {

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println("ERR msg:", err)
	}

}
