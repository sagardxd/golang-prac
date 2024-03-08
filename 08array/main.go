package main

import "fmt"

func main() {
	fmt.Println(("Welcome to array in golang"))

	var color [5]string

	color[0] = "pink"
	color[3] = "white"

	fmt.Println("color list is ", len(color))

	var race = [3]string{"white", "caucasion", "black"}
	fmt.Println("race list is ", race)

}
