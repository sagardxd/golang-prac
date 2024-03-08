package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var num float64 = 2
	var ptr = &num

	fmt.Println("value of ptr: ", ptr)
	fmt.Println("value inside ptr: ", *ptr)

	fmt.Println("give a input value you want to add in num:")


	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	inputValue, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	*ptr = *ptr + inputValue
	fmt.Println("Added: ", *ptr)

}
