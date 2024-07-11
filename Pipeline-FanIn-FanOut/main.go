package main

import "fmt"

func main() {
	randomNumbers := []int{}
	sum := 0

	for i := 1; i <= 100000; i++ {
		randomNumbers = append(randomNumbers, i)
	}

	//generate pipeline
	generatePipeline := generatePipeline(randomNumbers)

	//fan out
	c1 := fanOut(generatePipeline)
	c2 := fanOut(generatePipeline)
	c3 := fanOut(generatePipeline)
	c4 := fanOut(generatePipeline)
	c5 := fanOut(generatePipeline)
	c6 := fanOut(generatePipeline)
	c7 := fanOut(generatePipeline)
	c8 := fanOut(generatePipeline)
	c9 := fanOut(generatePipeline)

	//fan in
	c := fanInt(c1, c2, c3, c4, c5, c6, c7, c8, c9)

	for i := 0; i <= len(randomNumbers); i++ {
		sum += <-c
	}

	fmt.Println("Total: ", sum)

}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()

	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func fanInt(inputChannel ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for _, c := range inputChannel {
			for n := range c {
				out <- n
			}
		}
		close(out)
	}()
	return out
}
