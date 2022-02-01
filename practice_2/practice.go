package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	i1, i2, i3 := 12, 45, 68

	intSum := i1 + i2 + i3

	fmt.Println(intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3

	floatSum := f1 + f2 + f3

	floatSum = math.Round(floatSum*100) / 100
	//safe way to rounding fractional number

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("%.2f", circumference)

	n := time.Now()
	fmt.Println("Studied at", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Launched at", t)

	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("type of parsedTime %T", parsedTime)
}
