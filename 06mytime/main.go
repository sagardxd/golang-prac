package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("Present time is ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2003, time.May, 5, 5 ,0,0,0, time.UTC)
	fmt.Println("The time I created is: ", createdDate.Format("01-02-2006 15:04:05"));

}
