package main

import "fmt"

func main() {
	fmt.Println("Soumya")
	const DAILY_GOAL int = 10000
	steps := [7]int{4420, 5590, 7890, 2894, 8345, 4590, 9045}
	var totalSteps int
	for _, step := range steps {
		totalSteps += step
		if step < 5000 {
			fmt.Println("Too Low, try harder!")
		} else if step < DAILY_GOAL {
			fmt.Println("Good, but you can do more!")
		} else {
			fmt.Println("Great! Goal Achieved!")
		}
	}

	var averageSteps = totalSteps / 7
	fmt.Println(averageSteps)

	switch {
	case averageSteps <= 5000:
		fmt.Println("Poor Week")
	case averageSteps < 10000 && averageSteps >= 5000:
		fmt.Println("Good Week")
	case averageSteps >= 10000:
		fmt.Println("Excellent Week!")

	}
}
