package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var res string
		fmt.Printf("Recieved Message: %s\n", data)
		if string(data) == "ping" {
			res = "pong"
		} else {
			res = string(data)
		}
		if err := conn.WriteMessage(messageType, []byte(res)); err != nil {
			log.Println(err)
			return
		}
	}

}

func main() {
	http.HandleFunc("/ws", handleWs)

	fmt.Println("ws server started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
