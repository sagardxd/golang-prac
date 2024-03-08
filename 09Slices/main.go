package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi")

	var fruits = []string{"Tomato", "Apple", "Peach"}
	fruits = append(fruits, "mango")

	fmt.Println(fruits)

	fruits = append(fruits[1:4])
	fmt.Println(fruits)

	highScores := make([]int, 4)

	highScores[0] = 300
	highScores[1] = 900
	highScores[2] = 100
	highScores[3] = 200

	// fmt.Println(highScores)
	// highScores = append(highScores, 211, 1000, 345)
	// sort.Ints(highScores)
	// fmt.Println(highScores)

	//how to remove a particular index
	courses := []string{"react", "next", "tailwind", "ruby", "rust"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
