package main

import "fmt"

func main() {
	//while like
	const maxLimit = 5
	i := 1

	for i < maxLimit {
		fmt.Println(i)
		i = i + 1
	}

	for x := 10; x > 0; x-- {
		if x == 4 {
			continue
		}
		if x == 2 {
			break
		}
		fmt.Println(x)
	}

	const barAge int = 18
	const haveDL = true

	if barAge >= 18 && haveDL {
		fmt.Println("A rider ready to enter in bar")
	}

	const mark = 23

	if mark > 90 {
		fmt.Println("A+")
	} else if mark > 50 {
		fmt.Println("B")
	} else {
		fmt.Println("Try to score more")
	}
}
