package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Harman6282/blog-go/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", handlers.Greet)
	r.Get("/blogs", handlers.GetAllblogs)
	r.Post("/blogs", handlers.Createblog)
	r.Get("/blogs/{id}", handlers.GetblogById)
	r.Put("/blogs/{id}", handlers.Updateblog)
	r.Delete("/blogs/{id}", handlers.Deleteblog)

	fmt.Println("Server started at Port 3000")
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatal("Error on starting server")
	}

}
