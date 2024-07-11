package main

import "fmt"

func main() {
	number := 3
	switch {
	case number <= 3:
		fmt.Println("Nho hon hoac bang 3")
		//fallthrough
		break
	case number <= 5:
		fmt.Println("Nho hon hoac bang 5")
		break
	default:
		fmt.Println("Ko thuoc truong hop nao")
	}

	var i interface{} = "Hello"
	switch i.(type) {
	case int:
		fmt.Println("Kieu du lieu int")
		break
	case float64:
		fmt.Println("Kieu du lieu float64")
		break
	default:
		fmt.Println("Ko biet kieu du lieu")
	}

	m := map[string]int{
		"Yuh":  19,
		"Mike": 20,
	}

	if age, isExisted := m["Yuh"]; isExisted == true {
		fmt.Println(age)
	}
}
