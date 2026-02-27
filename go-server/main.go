package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello sir"))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to golang app on AWS"))
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
