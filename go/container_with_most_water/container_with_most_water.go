package main

import (
	"fmt"
)

func maxArea(height []int) int {
	sizeIn := len(height)
	i := 0
	j := sizeIn - 1
	maxAreaValue := -1
	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > maxAreaValue {
			maxAreaValue = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxAreaValue
}

func main() {
	input1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input1))
	input2 := []int{1, 1}
	fmt.Println(maxArea(input2))
	input1 = []int{1, 1, 1, 1, 1, 1}
	fmt.Println(maxArea(input1))
	input1 = []int{1, 5, 1, 1, 4, 3}
	fmt.Println(maxArea(input1))
}
