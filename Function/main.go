package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

// phương thức
func (g greeter) greet() {
	g.name = "Yuh"
	fmt.Println(g.greeting, g.name)

}

func main() {
	a := 10
	b := 20
	max := findMax(a, b)
	fmt.Println(*max)
	fmt.Println(a)
	fmt.Println(findSum(1, 2, 3, 4, 5))

	g := greeter{
		greeting: "Hello",
		name:     "Anh",
	}

	g.greet()
	fmt.Println(g.name)
}

func findSum(a ...int) (sum int) { //...int = []int
	for _, v := range a {
		sum = sum + v
	}
	return
}

// đối với kiểu dữ liệu ko phải slice hoặc map, khi chúng ta truyền vào tham số của hàm thì nó sẽ truyền vào bảng copy
// tức là tạo vùng nhớ mới, sau đó đưa giá trị tham số vào vùng nhớ mới đó, và vào thân hàm thì nó sẽ thao tác trong vùng nhớ mới
// thay đổi các giá trị tính toán thì ko thay đổi biến mà truyền vào hàm
func findMax(a, b int) *int {
	a = 100
	if a > b {
		return &a
	}
	return &b
}

func findMaxPointer(a, b *int) int {
	*a = 100
	if *a > *b {
		return *a
	}
	return *b
}
