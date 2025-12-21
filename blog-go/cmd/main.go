package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Harman6282/blog-go/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Greet).Methods("GET")
	r.HandleFunc("/blogs", handlers.GetAllblogs).Methods("GET")
	r.HandleFunc("/blogs", handlers.Createblog).Methods("POST")
	r.HandleFunc("/blogs/{id}", handlers.GetblogById).Methods("GET")
	r.HandleFunc("/blogs/{id}", handlers.Updateblog).Methods("PUT")
	r.HandleFunc("/blogs/{id}", handlers.Deleteblog).Methods("DELETE")

	fmt.Println("Server started at Port 3000")
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatal("Error on starting server")
	}

}
