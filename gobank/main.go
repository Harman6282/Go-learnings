package main

import "fmt"

func main(){
	server := NewApiServer(":8080")
	server.Run()
	fmt.Println("Heloo")
}
