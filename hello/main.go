package main

import (
	"fmt"
)

const constInt int = 1000

func main() {

	var aString string = "This is Go"

	fmt.Println(aString)
	fmt.Printf("types is %T\n", aString)

	var anInteger int = 42

	fmt.Println(anInteger)
	fmt.Printf("types is %T\n", anInteger)

	var defaultInt int

	fmt.Println(defaultInt)

	var anotherInt = 100

	fmt.Println(anotherInt)
	fmt.Printf("types is %T\n", anotherInt)

	//should be inside the function
	anotherString := "This is also String"

	fmt.Println(anotherString)
	fmt.Printf("types is %T\n", anotherString)

}
