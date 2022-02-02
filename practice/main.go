package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")

	//underscore: to ignore variable
	input, _ := reader.ReadString('\n')
	fmt.Println("Your input: ", input)	

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		fmt.Println(err)
		panic("value must be a number")
		//custom message and exit app
	} else {
		fmt.Println("value of number: ", aFloat)
	}

	var anInt int = 5
	var otherFloat float64 = 42
	sum := float64(anInt) + otherFloat
	fmt.Print(sum)

}
