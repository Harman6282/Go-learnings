package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Harman6282/blog-go/internal/types"
	"github.com/Harman6282/blog-go/internal/utils/response"
	"github.com/go-chi/chi/v5"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, http.StatusOK, "hello from server")
}

func GetAllblogs(w http.ResponseWriter, r *http.Request) {

	response.WriteJson(w, http.StatusOK, Blogs)

}

func Createblog(w http.ResponseWriter, r *http.Request) {
	var blog types.Blog

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if blog.Title == "" {
		response.WriteJson(w, http.StatusBadRequest, "Enter title field")
		return
	}

	response.WriteJson(w, http.StatusOK, blog)
	fmt.Println(blog)

}

func GetblogById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Println("id is: ", id)


}

func Updateblog(w http.ResponseWriter, r *http.Request) {

}

func Deleteblog(w http.ResponseWriter, r *http.Request) {

}
