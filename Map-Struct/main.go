package main

import "fmt"

type People struct {
	name   string
	age    int
	isMale bool
}

type Student struct {
	People
	number   int
	subjects []string
}

func main() {
	// Map
	// studentNameAgeMap := make(map[string]int)
	// studentNameAgeMap = map[string]int{
	// 	"Yuh":  19,
	// 	"Tom":  35,
	// 	"Mike": 40,
	// }
	// fmt.Println(studentNameAgeMap)

	// studentNameAgeMap["Yuh"] = 25
	// fmt.Println(studentNameAgeMap)

	// delete(studentNameAgeMap, "Mike")
	// fmt.Println(studentNameAgeMap)

	// _, isExisted := studentNameAgeMap["Yuh"]
	// fmt.Println(isExisted)

	// copyMap := studentNameAgeMap
	// fmt.Println(studentNameAgeMap)
	// copyMap["Yuh"] = 30
	// fmt.Println(studentNameAgeMap)
	// fmt.Println(copyMap)

	// Struct
	// studentYuh := Student{
	// 	number:   3,
	// 	name:     "Yuh",
	// 	isMale:   true,
	// 	subjects: []string{"Math", "English", "Computer"},
	// }

	studentYuh := Student{}
	studentYuh.name = "Yuh"
	studentYuh.age = 20
	studentYuh.number = 10
	studentYuh.isMale = true
	studentYuh.subjects = []string{"Math", "English", "Computer"}
	fmt.Println(studentYuh)

	studentYuh1 := struct{ name string }{name: "Yuh"}
	copyStudent := studentYuh1
	copyStudent.name = "Mike"

	fmt.Println(studentYuh1)
	fmt.Println(copyStudent)
}
