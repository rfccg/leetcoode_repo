package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ans := mergeArr(nums1, nums2)
	size := len(ans)
	fmt.Println(ans)
	if size%2 == 0 {
		return (float64(ans[size/2]) + float64(ans[-1+size/2])) / 2
	}

	return float64(ans[size/2])
}

func mergeArr(nums1 []int, nums2 []int) []int {
	newArr := make([]int, len(nums1)+len(nums2))
	idx1 := 0
	idx2 := 0
	for i := 0; i < len(newArr); i++ {
		sizable1 := idx1 < len(nums1)
		sizable2 := idx2 < len(nums2)

		if sizable1 && sizable2 && nums1[idx1] < nums2[idx2] {
			newArr[i] = nums1[idx1]
			idx1++
		} else if sizable1 && sizable2 {
			newArr[i] = nums2[idx2]
			idx2++
		} else if sizable1 {
			newArr[i] = nums1[idx1]
			idx1++
		} else {
			newArr[i] = nums2[idx2]
			idx2++
		}
	}
	return newArr
}

func main() {
	arr1 := []int{1, 4, 5}
	arr2 := []int{2, 3, 6, 7}

	ans := findMedianSortedArrays(arr1, arr2)
	fmt.Println(ans)
}
