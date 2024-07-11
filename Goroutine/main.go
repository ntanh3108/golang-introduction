package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.Mutex{}

func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(2)
		m.Lock()
		go sayHello()
		m.Lock()
		go add()
	}
	wg.Wait()
} //main cung la 1 goroutine

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.Unlock()
	wg.Done()
}

func add() {
	counter++
	m.Unlock()
	wg.Done()
}
