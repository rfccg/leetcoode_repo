package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	var newStr strings.Builder
	currLine := 0
	currPos := 0
	offset := numRows // i*len(s) = j
	for i := 0; i < len(s); i++ {
		if currPos >= len(s) {
			currLine++
			currPos = currLine
			offset = currPos + 2*(numRows-currLine) - 2
		}
		newStr.WriteByte(s[currPos])
		// fmt.Printf("%d, %d\n", currPos, offset)
		currPos += 2*numRows - 2
		if currLine > 0 && currLine < numRows-1 && offset < len(s) {
			// 0 new2
			i++
			newStr.WriteByte(s[offset])
			offset += 2*numRows - 2
		}

	}
	return newStr.String()
}

// 0,n-1

func main() {
	fmt.Println(-10 % 2)
	s := "PAYPALISHIRING"
	lineNum := 3
	ans := "PAHNAPLSIIGYIR"
	fmt.Println(convert(s, lineNum))
	fmt.Println(ans)

	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println("PINALSIGYAHRPI")

	fmt.Println(convert("abc", 2))
	fmt.Println("acb")

	fmt.Println(convert("abc", 1))
	fmt.Println("abc")
}

//
// 1   5   9
// 2 4 6 8 10
// 3   7
//
// 1     7
// 2   6 8
// 3 5   9
// 4     10
//
// 1 + 6 = 7
// 7 + 6 -> bleh
// 2
