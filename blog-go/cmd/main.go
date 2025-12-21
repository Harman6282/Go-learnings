package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Harman6282/blog-go/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Greet)
	r.HandleFunc("/blogs", handlers.GetAllblogs)
	r.HandleFunc("/blogs", handlers.Createblog)
	r.HandleFunc("/blogs/{id}", handlers.GetblogById)
	r.HandleFunc("/blogs/{id}", handlers.Updateblog)
	r.HandleFunc("/blogs/{id}", handlers.Deleteblog)

	fmt.Println("Server started at Port 3000")
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatal("Error on starting server")
	}

}
