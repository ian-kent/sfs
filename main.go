package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	bindAddr := ":4242"
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	log.Printf("Serving %s on %s", dir, bindAddr)
	err := http.ListenAndServe(bindAddr, http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Printf("Error binding to %s: %v", bindAddr, err)
	}
}
