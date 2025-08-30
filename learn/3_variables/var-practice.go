// package main

// import "fmt"

// func main() {
// 	var studentName string = "S Soumyakanta"
// 	fmt.Println(studentName)

// 	//no need to define type

// 	var studentMarksInEnglish = 50
// 	fmt.Println(studentMarksInEnglish)

// 	//shorthand for variable
// 	studentAge := 33
// 	fmt.Println(studentAge)

// 	var studentSection string
// 	studentSection = "B"
// 	fmt.Println(studentSection)
// }

package main

import "fmt"

func main() {
	var name string
	age := 99
	name = "Soumya"
	mob := 88
	fmt.Println(name, age, mob)

	var name2 string
	var age2 int
	var minor bool
	var height float32

	fmt.Println(name2, age2, height, minor)

	var yearOne, yearTwo int = 1894, 3294
	fmt.Println(yearTwo)
	fmt.Println((yearOne))
	// fmt.Println((yearOne)) works because Go allows unnecessary grouping parentheses.

	var temp int = 32
	_ = temp // no error, safely ignored using blank identifier
}
