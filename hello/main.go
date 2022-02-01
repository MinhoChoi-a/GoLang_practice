//Go programs start running in the main package. It is a special package that is used with programs that are meant to be executable
package main

import (
	"fmt"
)

const constInt int = 1000

//The main() function is a special function that is the entry point of an executable program.
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
