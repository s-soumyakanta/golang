package main

import (
	"fmt"
	"time"
)

func main() {
	//week name as per number

	i := 4

	switch i {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Not a weekday")
	}

	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work day")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {

		//we can also use it like this
		// switch i.(type) {

		case int:
			fmt.Println("Intiger")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI("S Soumyakanta")
	whoAmI(3.3)
	whoAmI(true)
}
