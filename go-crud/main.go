package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods(("GET"))
	router.HandleFunc("/users/{id}", getUser(db)).Methods(("GET"))
	router.HandleFunc("/users", createUser(db)).Methods(("POST"))
	router.HandleFunc("/users/{id}", updateUser(db)).Methods(("PUT"))
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods(("DELETE"))

	// start server
	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// get all users

func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(users)
	}
}

// get user by id

func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var u User
		err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.Id, &u.Name, &u.Email)
		if err != nil {
			// fix error handling
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

// create user

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1 , $2) RETURNING id", u.Name, u.Email)
		if err != nil {
			// fix error handling
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

func updateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE users SET name = $1, email = $2, WHERE id = $3", u.Name, u.Email, id)
		if err != nil {
			// fix error handling
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("DELETE * FROM users WHERE id = $1", id)
		if err != nil {
			// fix error handling
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode("User deleted")
	}
}
