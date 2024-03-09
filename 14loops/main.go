package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		// fmt.Printf("%v ", days[d])
	}

}
