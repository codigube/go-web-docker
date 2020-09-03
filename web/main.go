package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Simple Web Server
	http.HandleFunc("/", HelloServer)

	portNum := os.Getenv("PORT")
	if portNum == "" {
		portNum = "8080"
	}

	log.Printf("Exposed port: %v", portNum)

	http.ListenAndServe(":"+portNum, nil)
}

// HelloServer -
func HelloServer(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "Hello world from Go Web Server!")

	// Log
	log.Printf("GO Environment: %v", os.Getenv("GO_ENV"))
}
