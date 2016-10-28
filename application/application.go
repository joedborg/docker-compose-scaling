package main

import (
	"os"
	"io"
	"net/http"
	"log"
)

func index(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	io.WriteString(w, hostname)
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
