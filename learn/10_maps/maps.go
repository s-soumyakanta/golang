package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map - similar to object in JavaScript

	m := make(map[string]string)

	//setting an element

	m["name"] = "S Soumyakanta"
	m["language"] = "JavaScript"

	//get an element

	fmt.Println(m["name"], m["language"]) //S Soumyakanta JavaScript

	// while accessing non existing key
	// Imp: if key does not exists in the map then it returns zero value
	fmt.Println(m["phone"]) // we are getting empty string

	marks := make(map[string]int)

	marks["Math"] = 30
	marks["Science"] = 90

	fmt.Println(marks)      //map[Math:30 Science:90]
	fmt.Println(len(marks)) //2

	delete(marks, "Science")

	fmt.Println(len(marks)) //1

	clear(marks) // empty now

	//making maps without using make - if we know elements in advance
	phNums := map[string]int{"Soumya": 123, "SS": 454}

	fmt.Println(phNums) //map[SS:454 Soumya:123]

	//getting element in go
	k, ok := phNums["SS"]
	fmt.Println(k) //454
	if ok {
		fmt.Println("All ok")
	} else {
		fmt.Println("Not ok")
	}

	map1 := map[string]int{"price": 44, "phones": 3}
	map2 := map[string]int{"price": 44, "phones": 3}

	//check whethere two maps are equal or not

	fmt.Println(maps.Equal(map1, map2)) //true

}
