package main

import "fmt"

//only for in go no while -> only construct in go for looping
func main() {
	//while loop

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	//infinite loop
	// for {
	// 	fmt.Println("ss")
	// }

	//classic for loop

	for i := 20; i < 28; i++ {
		// break
		//break breack the code

		//continue
		if i == 25 {
			//continue skip the current itteration -> it skipped to pring number 30
			continue
		}
		fmt.Println(i + 5)
	}

	//range - used to a certain stuff for a time
	for i := range 5 {
		fmt.Println(i)
	}

}
