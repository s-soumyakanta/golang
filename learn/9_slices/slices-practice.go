package main

import "fmt"

func main() {
	names := [4]string{
		"Soumya",
		"Kanta",
		"Dev",
		"Name",
	}
	fmt.Println(names, len(names)) //[Soumya Kanta Dev Name] 4

	var ages = []int{2, 4, 9}
	ages[2] = 7
	// ages[3] = 19 panic: runtime error: index out of range [3] with length 3
	fmt.Println(cap(ages), ages) //3 [2 4 7]

	nums := make([]int, 2, 5)
	// nums = append(nums, 2, 3, 432, 342, 53) //7 [0 0 2 3 432 342 53] 10
	fmt.Println(len(nums), nums, cap(nums)) //2 [0 0] 5

	var randomArr []int

	randomArr = append(randomArr, 22)
	fmt.Println(randomArr)
	//copy
	firstCopy := make([]string, 0)
	firstCopy = append(firstCopy, "Soumya")
	secondCopy := make([]string, len(firstCopy))
	copy(secondCopy, firstCopy)
	fmt.Println(firstCopy, secondCopy) //[Soumya] [Soumya]
	//slice operator
	zips := make([]int, 0)
	zips = append(zips, 242, 973, 465, 783, 690)
	fmt.Println(zips) //[242 973 465]

	fmt.Println(zips[1:3]) //[973 465]
	fmt.Println(zips[:2])  //[242 973]
	fmt.Println(zips[1:])  //[973 465 783 690]
}
