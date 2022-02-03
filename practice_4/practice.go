package main

import (
	"fmt"
)

//using map: unordered data structure, using key and value

func main() {

	//make() initializes memory, but new() does not. It just gives a name of variable
	states := make(map[string]string)
	fmt.Println(states)

	states["wa"] = "washington"
	states["or"] = "oregon"

	fmt.Println(states)

	washinton := states["wa"]

	fmt.Println(washinton)

	delete(states, "or")
	states["ny"] = "new york"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))

	i := 0

	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)

	//go doesn't have while
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

	colors := []string{"a", "b", "c"}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1

	for value < 10 {
		fmt.Println(value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum

		fmt.Println("sum")

		if sum < 200 {
			goto theEnd
		}
	}

	//label
theEnd:
	fmt.Println("End")
}
