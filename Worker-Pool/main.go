package main

import "fmt"

func main() {
	numbers := 100000
	numberOfWorkers := 10
	jobs := make(chan int, numbers)
	results := make(chan int, numbers)

	for i := 0; i <= numberOfWorkers; i++ {
		go worker(jobs, results)
	}

	for i := 0; i <= numbers; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i <= numbers; i++ {
		fmt.Println(<-results)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- Fibonacci(n)
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
