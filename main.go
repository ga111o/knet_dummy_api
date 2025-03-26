package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		fmt.Fprintf(w, `{"Success": true, "Detail": "Hello, This is K-NET"}`)
	})

	fmt.Println("Server starting on :60950")
	if err := http.ListenAndServe(":60950", nil); err != nil {
		log.Fatal(err)
	}
}
