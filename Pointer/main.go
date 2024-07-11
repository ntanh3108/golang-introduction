package main

import "fmt"

type myStruct struct {
	number int
}

func main() {
	var a int = 12
	var b *int = &a
	fmt.Println(a, *b)

	a = 24
	fmt.Println(a, *b)

	c := [3]int{1, 2, 3}
	d := c
	fmt.Println(c, d)

	c[1] = 9
	fmt.Println(c, d)

	e := []int{4, 5, 6}
	f := e
	fmt.Println(e, f)

	e[1] = 9
	fmt.Println(e, f)

	var g *myStruct
	fmt.Println(g)
	g = new(myStruct)
	fmt.Println(g)
	g.number = 20
	fmt.Println(g.number)
}

/*Slice trong golang nó chứa 2 con trỏ, phía bên dưới slice là array. Slice a chứa 2 con trỏ, con trỏ 1 trỏ tới phần tử đầu,
con trỏ 2 trỏ tới phần tử cuối. Khi khai báo slice b = a, thì t đang khai báo con trỏ b trỏ tới a */
/* map tương tự slice */
