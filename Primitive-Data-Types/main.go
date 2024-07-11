package main

import (
	"fmt"
)

func main() {
	var a bool
	a = 1 == 1

	var b int8 = -32

	c := 13.7e72

	var d complex64 = 1 + 2i

	var e string = "Hello"

	var f []byte = []byte(e)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", real(d), real(d))
	fmt.Printf("%v, %T\n", imag(d), imag(d))
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)
}
