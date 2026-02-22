package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, resuls chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range jobs {
		fmt.Println("image", i, "processing by worker ", id)
		time.Sleep(time.Second)
		fmt.Println("worker-", id, "finished job ", i)
		resuls <- i * 2
	}
}

func main() {
	var wg sync.WaitGroup

	numJobs := 6
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

}
