package main

import "fmt"

func longestPalindrome(s string) string {
	longestSize := 1
	startIndex := 0
	for i := 0; i < len(s); i++ {
		currSize := 1
		currStart := i
		evenOk := true
		oddOk := true
		for j := 1; j <= (i + 1); j++ {

			lowerBoundCheck := (i - j) >= 0
			upperBoundEven := (i + j - 1) < len(s)
			upperBoundOdd := (i + j) < len(s)
			if lowerBoundCheck && upperBoundEven && s[i+j-1] == s[i-j] && evenOk { // even size
				currStart = i - j
				currSize = 2 * j
			} else {
				evenOk = false
			}
			if lowerBoundCheck && upperBoundOdd && s[i+j] == s[i-j] && oddOk { // odd size
				currStart = i - j
				currSize = j*2 + 1
			} else {
				oddOk = false
			}
			if !evenOk && !oddOk {
				break
			}
		}

		if longestSize < currSize {
			longestSize = currSize
			startIndex = currStart
		}
	}
	//
	return s[startIndex:(startIndex + longestSize)]
}

func main() {
	test := "aaa"
	ans := longestPalindrome(test)
	fmt.Println(ans == test)

	test = "aa"
	ans = longestPalindrome(test)
	fmt.Println(ans == test)

	test = "aaaa"
	ans = longestPalindrome(test)
	fmt.Println(ans == test)

	test = "iasa"
	ans = longestPalindrome(test)
	fmt.Println(ans == "asa")
}
