package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan int, 50)
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)

	wg.Add(2)
	// go func(ch <-chan int) { //nhan
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) { //gui
	// 	ch <- 42
	// 	wg.Done()
	// }(ch)

	// go func(ch <-chan int) { //nhan
	// 	for {
	// 		if i, ok := <-ch; ok {
	// 			fmt.Println(i)
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) { //gui
	// 	ch <- 42
	// 	ch <- 42
	// 	close(ch)
	// 	wg.Done()
	// }(ch)

	go func() { //nhan
		for {
			select {
			case i := <-ch1:
				fmt.Printf("Channel 1: %v\n", i)
			case j := <-ch2:
				fmt.Printf("Channel 2: %v\n", j)
			default:
				z := <-ch
				fmt.Printf("Channel 2: %v\n", z)
			}
		}
		wg.Done()
	}()

	go func() { //gui
		ch1 <- 42
		close(ch1)
		ch2 <- 50
		close(ch2)
		wg.Done()
	}()
	wg.Wait()
}
