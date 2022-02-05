package main

import (
	"fmt"
)

func main() {

	sum := addValue(5, 8)

	fmt.Println(sum)

	multiSum, count := addAll(4, 7, 9)

	//defer keyword, function call is not executed until the surrounding functions returns.
	defer fmt.Println(multiSum)
	fmt.Println(count)

}

func addValue(v1, v2 int) int {
	return v1 + v2

}

func addAll(values ...int) (int, int) {
	total := 0

	//don't care index
	for _, v := range values {
		total += v
	}

	return total, len(values)
}
