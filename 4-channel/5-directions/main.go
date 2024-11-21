package main

import "fmt"

// Thread - 1
func main() {
	chn := make(chan string)
	go receive("Diego", chn)
	read(chn)
}

func receive(name string, chn chan<- string) {
	chn <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}
