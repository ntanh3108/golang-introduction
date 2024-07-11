package main

import "fmt"

var (
	n      int    = 10
	m      int    = 20
	Number int    = 500 //export with first Uppercase
	h      string = "abc"
)

func main() {
	// var i int
	// i = 2024
	// var i int = 10
	i := 3000

	fmt.Println(n)
	var n int = 1000 //shadow in block scope

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(h)
	fmt.Printf("%v, %T", i, i)
}

/*
1. Dinh nghia bien trong golang
2. Cu phap khai bao bien
3. Global and block scope
4. Shadow
5. Declared and not used
6. Export global scope
7. Naming convention -> camelCase
8. Convert type
*/
