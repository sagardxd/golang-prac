package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome nigga"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate my pp")

	//comma ok , comma err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of this rating: %T", input)

}
