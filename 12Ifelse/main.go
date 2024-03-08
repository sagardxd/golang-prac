package main

import "fmt"

func main() {
	fmt.Println("if else in golang")
	isLogin := false

	if isLogin {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if num := 3; num < 10 {
		fmt.Println("less than three")
	}
}
