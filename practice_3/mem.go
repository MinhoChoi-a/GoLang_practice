package main

import (
	"fmt"
	"sort"
)

func main() {

	//just set a pointer p, data is nil(null)
	//it doesn't contain anything
	var p *int

	fmt.Println("Value of p:", p)

	//fmt.Println("Value of p:", *p)
	// *p : pointing some other variable
	//will get error, because p is pointing at an invalid memory address (isn't pointing anything)

	anInt := 42
	var pp = &anInt
	//&(ampersand): pointing the memory address of the variable, x value
	fmt.Println("Value of p:", *pp) //is working

	*pp = *pp * 2
	fmt.Println("Value of p:", anInt) //is working

	//same concept with reference in Java
	//but unlike with Java, it doesn't have to point some particular value initially, and change it in runtime

	//array
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"

	fmt.Println(colors)

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(len(numbers))

	//slices: abstraction layer that sits on top of an array
	//re-sizeable, can be sorted easily

	//just need to remove the length of array
	var slice = []string{"a", "B", "c"}

	slice = append(slice, "d")
	fmt.Println(slice)

	slice = append(slice[1:len(slice)])
	fmt.Println(slice)

	slice = append(slice[:len(slice)-1])
	fmt.Println(slice)

	makeSlice := make([]int, 5, 5) //type, size, capacity(limit of size)
	makeSlice[0] = 1
	makeSlice[1] = 4
	makeSlice[2] = 2
	makeSlice[3] = 3
	makeSlice[4] = 5

	sort.Ints(makeSlice)
	fmt.Println(makeSlice)
}
