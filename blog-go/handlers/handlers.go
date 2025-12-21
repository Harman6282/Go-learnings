package handlers

import (
	"net/http"

	"github.com/Harman6282/blog-go/utils/response"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, http.StatusOK, "hello from server")
}

func GetAllblogs(w http.ResponseWriter, r *http.Request) {

	response.WriteJson(w, http.StatusOK, Blogs)

}

func Createblog(w http.ResponseWriter, r *http.Request) {

}

func GetblogById(w http.ResponseWriter, r *http.Request) {

}

func Updateblog(w http.ResponseWriter, r *http.Request) {

}

func Deleteblog(w http.ResponseWriter, r *http.Request) {

}
