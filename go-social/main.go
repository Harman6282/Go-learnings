package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func dbConnect () (*sql.DB, err error){
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		log.Fatal("Check your creadentials in evn")
	}
} 

type Post struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

func getAllposts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go server"))
}

func createpost(w http.ResponseWriter, r *http.Request) {

}

func getpostById(w http.ResponseWriter, r *http.Request) {

}

func updatepost(w http.ResponseWriter, r *http.Request) {

}

func deletepost(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/posts", getAllposts)
	r.HandleFunc("/posts", createpost)
	r.HandleFunc("/posts/{id}", getpostById)
	r.HandleFunc("/posts/{id}", updatepost)
	r.HandleFunc("/posts/{id}", deletepost)

	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
