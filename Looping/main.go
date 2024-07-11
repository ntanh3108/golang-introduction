package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// array := [3]int{2, 3, 4}
	// for i, v := range array {
	// 	fmt.Println(i, v)
	// }

	m := map[string]int{
		"Yuh":  19,
		"Mike": 20,
		"Jack": 5,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
