package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)

	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			i := <-messages
			fmt.Printf("send message: %d\n", i)
		}
	}()

	// producer
	go func() {
		for i := 0; i < 12; i++ {
			messages <- i
			fmt.Printf("produced: %d\n", i)
		}
	}()
	time.Sleep(15 * time.Second)
	fmt.Println("main process exit!")
}
