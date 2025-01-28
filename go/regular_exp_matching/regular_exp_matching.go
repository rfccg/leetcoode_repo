package main

import "fmt"

func isMatchReverse(s string, p string, wildcard bool) bool {
	idxP := len(p) - 1
	idxS := len(s) - 1

	if idxS >= 0 && idxP < 0 {
		return false
	}
	if idxS < 0 && idxP < 0 {
		return true
	}
	if p[idxP] == byte('*') {
		return isMatchReverse(s, p[:idxP], true)
	}

	hasWildcard := false
	if wildcard {
		hasWildcard = isMatchReverse(s, p[:idxP], false)
	}
	if idxS < 0 {
		return hasWildcard
	}
	if s[idxS] == p[idxP] || p[idxP] == byte('.') {
		if wildcard {
			hasWildcard = hasWildcard || isMatchReverse(s[:idxS], p, true)
		}

		return hasWildcard || isMatchReverse(s[:idxS], p[:idxP], false)
	}

	return hasWildcard
}

func isMatch(s string, p string) bool {
	return isMatchReverse(s, p, false)
}

func main() {
	a := "123456"
	fmt.Println(a[:5])
	fmt.Println(isMatch("aa", "a"))             // false
	fmt.Println(isMatch("aa", "a*"))            // true
	fmt.Println(isMatch("ab", ".*"))            // true
	fmt.Println(isMatch("aa", "ab"))            // false
	fmt.Println(isMatch("aax", ".b"))           // false
	fmt.Println(isMatch("aa", "b"))             // false
	fmt.Println(isMatch("aab", "c*a*b"))        // true
	fmt.Println(isMatch("aasdxssxa", "a*.xx"))  // false
	fmt.Println(isMatch("aasdxssxa", "a*.x."))  // false
	fmt.Println(isMatch("aasdxssxa", "a.*.xa")) // true
	fmt.Println(isMatch("aaa", "ab*a*c*a"))     // true
	fmt.Println(isMatch("a", "ab*"))            // true
	fmt.Println(isMatch("a", "ab*"))            // true
	fmt.Println(isMatch("bbbba", ".*a*a"))      // true
	fmt.Println(isMatch("a", ".*"))             // true
	fmt.Println(isMatch("ab", ".*c"))           // false
}
