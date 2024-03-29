package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	theAnswer := 42
	var result string

	//opening brace much be in the same line with if statement
	if theAnswer < 0 {
		result = "Less then zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less then zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	switch dow {
	case 1:
		result = "Weekend"
		//fallthrough: will not go through the next case
	default:

		result = "weekday"
	}

	fmt.Println(result)
}
