package main

import (
	"fmt"
	// "math"
)

func reverse(x int32) int32 {
	acc := int32(0)
	multiplier := int32(10)
	overflowDetection := acc
	for x != 0 {
		overflowDetection = acc
		acc *= multiplier
		digit := x % 10
		acc += digit
		x /= 10
		fmt.Printf("acc %d, ofd %d\n", acc, overflowDetection)
		if acc%10 != digit {
			// fmt.Println("Hit overflowDetection")
			// fmt.Printf("%d, %d \n", x, acc)
			return 0
		}
	}
	// maxValue := int32(math.Pow(2, 31))
	// fmt.Printf("%d\n", maxValue)
	// if acc > maxValue-1 || acc < -maxValue {
	// 	return 0
	// }
	return acc
}

func main() {
	a := int32(10)

	// fmt.Println(a * math.MaxInt32)
	// fmt.Println(reverse(10))
	// fmt.Println(reverse(11))
	fmt.Println(int32(964632435) * a)
	fmt.Println(reverse(1534236469))
	// fmt.Println(reverse(210))
	// fmt.Println(reverse(101))
	// fmt.Println(reverse(-1315))
	// fmt.Println(reverse(1))
	// fmt.Println(reverse(0))
}
