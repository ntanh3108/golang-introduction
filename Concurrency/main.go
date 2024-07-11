package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 20; i++ {
			ch <- i
			fmt.Println("Sent ", i)
		}
	}()

	for i := 1; i <= 20; i++ {
		time.Sleep(time.Second)
		fmt.Println(<-ch)
	}
}
