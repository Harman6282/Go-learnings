package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dowork(done <-chan interface{}) {
	for {
		select {
			case <-done:
				return
		default:
			fmt.Println(rand.Int())
			time.Sleep(time.Millisecond)
		}

	}
}

func main() {

	done := make(chan interface{})

	go dowork(done)

	time.Sleep(time.Second * 2)
	close(done)
	
}
