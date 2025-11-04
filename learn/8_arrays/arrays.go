package main

import "fmt"

func main() {

	frnds := [4]string{"x", "y", "z"}
	frnds[3] = "o"
	fmt.Println(frnds[3])
	//2d arrays
	pins := [2][2]int{{753023, 234323}, {234523, 923423}}

	fmt.Println(pins) //[[753023 234323] [234523 923423]]

	//add values to array while declaring array or in single line

	ages := [4]int{3, 4, 2}

	fmt.Println(ages) //[3 4 2 0]

	//In golang arrays are number sequece of specific length -
	//we define number of element and they will be same type

	//we use arrays when their size is predictable or fixed size
	//memory optimization
	//constant time access of elements

	//in int type array by default for each element zero added
	var nums [5]int

	nums[0] = 33
	// empty string  added to array
	var names [3]string

	// falsy value added by default for array of bool
	var boolVals [4]bool
	boolVals[3] = true
	fmt.Println(boolVals) //[false false false true]

	fmt.Println(names) //[ ]

	fmt.Println(nums[0])   //33
	fmt.Println(nums)      //[33 0 0 0 0]
	fmt.Println(len(nums)) //5

}
