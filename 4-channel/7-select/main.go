package main

import (
	"fmt"
	"sync/atomic"
)

type Message struct {
	ID  int64
	Msg string
}

// Thread - 1
func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			c2 <- Message{i, "RabbitMQ message"}
		}

	}()
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			c2 <- Message{i, "Kafka message"}
		}

	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("received from rabbitMQ: ID: %d - %s\n", msg.ID, msg.Msg)
		case msg := <-c2:
			fmt.Printf("received from kafka: ID: %d - %s\n", msg.ID, msg.Msg)
		}
	}
}
