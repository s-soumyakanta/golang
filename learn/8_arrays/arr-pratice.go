// package main

// import "fmt"

// func main() {
// 	var arr1 = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(arr1)
// 	var arr2 = [...]string{"a", "b", "c", "d", "e"}

// 	arr2[4] = "f"

// 	arr3 := [3]int{1: 23}
// 	fmt.Println(arr2)
// 	fmt.Println(len(arr2))
// 	fmt.Println(arr3)
// 	//[0 23 0]
// }

package main

import "fmt"

func main() {
	var names = [3]string{"Soumya", "Sangram", "Subham"}
	fmt.Println(names[1])
	age := [...]int{3, 54, 3}
	age[2] = 43
	fmt.Println(age[2], len(age))

	const SUBJECTS int = 5

	var marks = [2][2]int{}

	for i := range marks {
		fmt.Println(i)
	}
}
