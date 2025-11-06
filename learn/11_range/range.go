package main

import (
	"fmt"
)

func main() {
	//we use range to iterating over data structures

	numbers := []int{4, 2, 4}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i]) // 4, 2, 4
	}

	//for loop with range

	sum := 0
	for ind, numbers := range numbers {
		fmt.Println(numbers) // 4, 2, 4

		sum = sum + numbers

		fmt.Println(numbers, ind) // 4 0 , 2 1, 4 2
	}

	fmt.Println(sum) // 10

	customers := map[string]string{"fName": "S", "lName": "Soumyakanta"}

	for k, v := range customers {
		fmt.Println(k, v) //fName S lName Soumyakanta
	}

	//using range over strings

	for i, c := range "golang" {
		fmt.Println(i, c)
		/* 0 103
		1 111
		2 108
		3 97
		4 110
		5 103
		*/
		// unicode code point rune
		// above the i is the starting byte of rune
		/* Multi-line comments start with /* and ends with */

		fmt.Println(i, string(c))
		/*0 g
		1 o
		2 l
		3 a
		4 n
		5 g */
	}
	weekend := make(map[int]string)
	weekend[1] = "saturday"
	weekend[2] = "sunday"
	for i, v := range weekend {
		fmt.Println(i, v) //1 saturday 2 sunday

	}
}
