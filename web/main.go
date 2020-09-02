package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Simple Web Server
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer -
func HelloServer(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "Hello world from Go Web Server!")
}
