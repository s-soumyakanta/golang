package main

import "fmt"

// we get a by value of argument, copy

//If we want to use the memory location of variable then we need to pass the variables ref
// * is pointer for memory ref of variable
func changeNum2(num *int) {
	// *num is derefference - the pointer/address memory of variable
	*num = 5
	fmt.Println(*num) //5

}
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num) //In changeNum 5
}

func changeSlice(nums []int) {
	nums = append(nums, 7)
}

// To change the values of slice we need to get it via ref, so we need to pass pointer
func changeSlice2(nums *[]int) {
	//we need to deref
	//*nums here we used it as it to change the value of pointer not the pointer
	*nums = append(*nums, 7)
}
func main() {
	//Pointers are the location of data saved in system
	num := 2
	num2 := 4
	changeNum(num)
	changeNum2(&num)
	fmt.Println("Memory address", &num)          //Memory address 0xc00000a0e8
	fmt.Println("Memory address", &num2)         //Memory address 0xc00000a100
	fmt.Println("After changeNum in main", num)  //After changeNum in main 2
	fmt.Println("After changeNum in main", num2) //After changeNum in main 4

	var nums = []int{3, 4, 9}
	var nums2 = []int{3, 4, 9}

	changeSlice(nums)
	//by value passed not by ref so 7 is not appended
	fmt.Println("After change", nums) //After change [3 4 9]

	//to get pointer we use & so &nums2
	changeSlice2(&nums2)

	fmt.Println(nums2) //[3 4 9 7]
	//7 is appended because of passing the memory ref using ref
}
