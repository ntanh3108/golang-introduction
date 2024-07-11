package main

import (
	"fmt"
)

const a = 42
const (
	red = iota
	yellow
	blue
	black
)

func main() {
	const i int16 = 42

	fmt.Printf("%v, %T\n", a, a)

	const a = 12
	fmt.Printf("%v, %T\n", a, a)

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", yellow, yellow)
	fmt.Printf("%v, %T\n", blue, blue)
	fmt.Printf("%v, %T\n", black, black)
}

/*
1. Dinh nghia hang so (constant)
2. Khai bao hang so
3. Dac diem cua hang so
4. Kieu du lieu Enum
5. Tu dong khoi tao gia tri cho Enum bang iota
*/
