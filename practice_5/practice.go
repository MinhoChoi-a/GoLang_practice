package main

import (
	"fmt"
)

//struct : is like the encapsulated class in Java -> but go doesn't have inheritance model(so it doesn't have any super/sub)

func main() {

	poodle := Dog{"Poodle", 10, "auf"}
	fmt.Println(poodle)

	//for debuggin process
	fmt.Printf("%+v\n", poodle)

	fmt.Printf("Breed: %v\n", poodle.Breed)
	poodle.Breed = "poddle"
	fmt.Printf("Breed: %v\n", poodle.Breed)

}

///uppercase initial character => exported and public
//lowercase initial character => non exported and private
// exported type needs a comment
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

//pass receiver(class)
//this is a brand new copy of the object.
//if you want to use it as Refernce, need to pass a pointer
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}
