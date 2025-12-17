package main
asdf
import(
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"json"
	"net/http"
)



func main(){
   
  r := mux.NewRouter()

  r.HandleFunc("/posts", getAllposts)
  r.HandleFunc("/posts", createpost)
  r.HandleFunc("/posts/{id}", getpostById)
  r.HandleFunc("/posts/{id}", updatepost)
  r.HandleFunc("/posts/{id}", deletepost)


}