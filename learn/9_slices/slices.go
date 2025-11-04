package main

import (
	"fmt"
	"slices"
)

func main() {
	// No need to specify the length of arrays. Dynamic arrays.
	// Most used construct in go
	// + useful methods

	//uninitialized slice is nil
	var nums []int
	fmt.Println(nums)        //currently it is empty
	fmt.Println(nums == nil) //true
	fmt.Println(len(nums))   //0

	var ages = make([]int, 2)
	//initial capacity is 2
	fmt.Println(cap(ages)) //2
	fmt.Println(ages)      // [0 0]

	var pins = make([]int, 2, 5)
	pins = append(pins, 2)
	pins = append(pins, 3)
	pins = append(pins, 4)
	pins = append(pins, 5)
	pins = append(pins, 6)
	pins = append(pins, 7)
	fmt.Println(pins)      //[0 0 2 3 4 5 6 7]
	fmt.Println(cap(pins)) //10 - on addition it doubles the capacity
	fmt.Println(len(pins)) //8

	//direct declare
	allNums := []int{}
	allNums = append(allNums, 2)
	fmt.Println(allNums)      //[2]
	fmt.Println(cap(allNums)) //1

	//adding elements using index
	var mobile = make([]int, 2, 5)

	mobile[0] = 34
	mobile[1] = 2
	fmt.Println(mobile) // [34 2]

	//copy function slice

	var studentsRollNumber = make([]int, 0, 5)
	studentsRollNumber = append(studentsRollNumber, 2)
	var studentsRollNumber2 = make([]int, len(studentsRollNumber))

	copy(studentsRollNumber2, studentsRollNumber)
	fmt.Println(studentsRollNumber, studentsRollNumber2) //[2] [2]

	//slice operator
	var allDigits = []int{1, 2, 3}

	fmt.Println(allDigits[0:2]) // [1,2]
	fmt.Println(allDigits[:1])  //[1]
	fmt.Println(allDigits[2:])  // [3]

	//slice package

	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}
	var nums3 = []int{1, 2, 3}

	fmt.Println(slices.Equal(nums1, nums2)) //true
	fmt.Println(slices.Equal(nums2, nums3)) //false

	//2d slices

	var nums4 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums4) //[[1 2 3] [4 5 6]]

}
