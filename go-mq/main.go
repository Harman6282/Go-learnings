package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type Queue struct {
	name string
	ch   chan string
}

// create queue
// publish to queue
// consume from queue

var (
	queues = make(map[string]*Queue)
	mu     sync.RWMutex
)

func createQueue(name string, size int) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exist := queues[name]; exist {
		return errors.New("queue already exists")
	}

	mq := &Queue{name: name, ch: make(chan string, size)}
	queues[name] = mq

	return nil
}

func getQueue(name string) (*Queue, bool) {
	mu.RLock()
	defer mu.RUnlock()
	mq, ok := queues[name]
	return mq, ok
}

func closeQueue(name string) error {
	mq, ok := getQueue(name)
	if !ok {
		return errors.New("queue not found")
	}
	close(mq.ch)
	return nil
}

// publish message to the queue
func publish(name, msg string) error {
	mq, ok := getQueue(name)
	if !ok {
		return errors.New("queue not found")
	}

	mq.ch <- msg
	return nil
}

// consume message from a queue
func consume(name string) (string, error) {
	mq, ok := getQueue(name)
	if !ok {
		return "", errors.New("queue not found")
	}
	msg, ok := <-mq.ch
	if !ok {
		return "", errors.New("queue closed")
	}
	return msg, nil
}

func main() {
	queueName := "chat"
	messages := []string{
		"Hello!",
		"How are you?",
		"I'm doing great, thanks",
		"Are we still meeting today?",
		"Yes, see you at 5",
		"Running a bit late",
		"No problem, take your time",
		"Can you send me the file?",
		"Just sent it",
		"Thanks, got it!",
	}

	if err := createQueue(queueName, 100); err != nil {
		log.Println(err)
	}
	for _, v := range messages {
		if err := publish(queueName, v); err != nil {
			log.Println(err)
		}
	}

	if err := closeQueue(queueName); err != nil {
		log.Println(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			msg, err := consume(queueName)
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(300 * time.Millisecond)
			fmt.Println(msg)
		}
	}()

	wg.Wait()

}
