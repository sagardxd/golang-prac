package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	sagar := User{"Sagar", "sagar@go", true, 20}

	fmt.Println(sagar)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
