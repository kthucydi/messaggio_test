package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("can not run listenandserve 8080: %v", err)
	}
}
