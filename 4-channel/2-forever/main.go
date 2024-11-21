package main

import "fmt"

// Thread - 1
func main() {
	forever := make(chan bool) // emtpy channel
	//Thread - 2
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		forever <- true // Fill channel
	}()

	<-forever
}
