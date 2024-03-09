package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		// fmt.Printf("%v ", days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	//returning the index and value
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index ,day);
	}

}
