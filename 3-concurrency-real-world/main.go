package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.lock()
		number++
		//atomic.AddUnint64(&number, 1) - this line will add lock and unlock for this part of the code
		//defer m.unlock()
		time.Sleep(1 * time.Second)
		w.Write([]byte(fmt.Sprintf("You are visitant number: %v", number)))
	})
	http.ListenAndServe(":3000", nil)
}
