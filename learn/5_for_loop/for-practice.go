package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println("Hello", i)
		i = i + 1
	}

	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	//reverse for loop

	for i := 10; i >= 0; i-- {
		if i == 5 {
			continue
		}
		fmt.Println(i)

	}

	//it will do till 4
	for i := range 5 {
		fmt.Println(i)
	}
}
