package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	acc := 0
	multiplier := 10
	overflowDetection := x > 0
	for x != 0 {
		acc *= multiplier
		digit := x % 10
		acc += digit
		x /= 10
		if (acc > 0) != overflowDetection && acc != 0 {
			// fmt.Println("Hit overflowDetection")
			// fmt.Printf("%d, %d \n", x, acc)
			return 0
		}
	}
	maxValue := int(math.Pow(2, 31))
	// fmt.Printf("%d\n", maxValue)
	if acc > maxValue-1 || acc < -maxValue {
		return 0
	}
	return acc
}

func main() {
	fmt.Println(reverse(10))
	fmt.Println(reverse(11))

	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(210))
	fmt.Println(reverse(101))
	fmt.Println(reverse(-1315))
	fmt.Println(reverse(1))
	fmt.Println(reverse(0))
}
