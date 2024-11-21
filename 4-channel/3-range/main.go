package main

import (
	"fmt"
)

// Thread - 1
func main() {
	ch := make(chan int)
	// Thread - 2
	go publish(ch)
	reader(ch)

}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Publishing %d\n", i)
		ch <- i
	}
	close(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
