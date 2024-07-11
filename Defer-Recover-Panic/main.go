package main

import "fmt"

func main() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
}

/* defer sẽ chứa trong ngăn xếp */
