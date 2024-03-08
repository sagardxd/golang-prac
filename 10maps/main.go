package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]int)

	languages["react"] = 2
	languages["tailwind"] = 2
	languages["js"] = 3
	languages["nodejs"] = 3

	fmt.Println(languages)
	fmt.Println("Score of react is", languages["react"])

	delete(languages, "tailwind")
	fmt.Println(languages)

	//looping in map

	for k, v := range languages {
		fmt.Printf("Key: %v and Value: %v \n", k,v)
	}
}
