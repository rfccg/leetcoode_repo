package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	chars := make(map[byte]int)
	startIdx := 0
	size := 0
	for i := 0; i < len(s); i++ {
		idx, ok := chars[s[i]]
		if ok {
			size = max(size, i-startIdx)
			if idx >= startIdx {
				startIdx = max(startIdx+1, idx+1)
			}
			fmt.Println(startIdx)
			fmt.Println(string(s[startIdx]))
		}
		chars[s[i]] = i

	}
	fmt.Println("------------")
	fmt.Println(len(s))
	fmt.Println(startIdx)
	fmt.Println(string(s[startIdx]))
	size = max(size, len(s)-startIdx)

	return size
}

func main() {
	test := "aaiajicde"
	fmt.Println(lengthOfLongestSubstring(test))
}
