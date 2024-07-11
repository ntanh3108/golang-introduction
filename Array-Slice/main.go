package main

import "fmt"

func main() {
	// array
	// var points [3]int
	// points := [3]int{7, 10, 5}
	//points := [...]int{7, 10, 5}

	//slice
	// a := []int{10, 20, 30, 40, 50, 60, 70}
	// b := a
	// b[1] = 50

	// c := a[:]
	// d := a[3:]
	// e := a[:6]
	// f := a[3:6]

	// fmt.Printf("%v, %v, %v\n", a, len(a), cap(a))
	// fmt.Printf("%v, %v, %v\n", b, len(b), cap(b))
	// fmt.Printf("%v, %v, %v\n", c, len(c), cap(c))
	// fmt.Printf("%v, %v, %v\n", d, len(d), cap(d))
	// fmt.Printf("%v, %v, %v\n", e, len(e), cap(e))
	// fmt.Printf("%v, %v, %v\n", f, len(f), cap(f))

	a := make([]int, 0)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	a = append(a, 1)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
	a = append(a, []int{2, 3, 4, 5, 6, 7}...)
	fmt.Printf("a: %v, %v, %v\n", a, len(a), cap(a))
}
