package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes Written: %v", n)
	})

	// Start a webserver
	_ = http.ListenAndServe(":8081", nil)
}
