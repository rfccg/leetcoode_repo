package main

import (
	"fmt"
	"sort"
)

type IdxSlice struct {
	sort.IntSlice
	idx []int
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func (s IdxSlice) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func twoSum(nums []int, target int) []int {
	idxArr := makeRange(0, len(nums))
	idxPreserve := IdxSlice{nums, idxArr}
	sort.Sort(idxPreserve)

	size := len(nums)
	i := 0
	j := size - 1

	for i < j {
		sum := idxPreserve.IntSlice[i] + idxPreserve.IntSlice[j]
		if sum == target {
			return []int{idxPreserve.idx[i], idxPreserve.idx[j]}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{idxPreserve.idx[i], idxPreserve.idx[j]}
}

func main() {
	numbers := []int{1, 2, 54, 7, 2, 3, 7, 8}
	twoSumR := twoSum(numbers, 15)

	fmt.Println(twoSumR)
}
