package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Addr string
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi there"))
}

func main() {

	r := chi.NewRouter()

	app := Application{
		Addr: ":8080",
	}

	r.Get("/", app.home)

	log.Println("server started")
	http.ListenAndServe(app.Addr, r)

}
