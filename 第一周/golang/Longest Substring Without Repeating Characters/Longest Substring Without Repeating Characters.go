package main

import (
	"fmt"
)

func main()  {
	k := "abcabcabc"
	fmt.Println(lengthOfLongestSubstring(k))
}

func lengthOfLongestSubstring(s string) int {
	var charpos [256]int
	start,max := 0,0
	for i := range charpos{
		charpos[i] = -1
	}
	for i := 0; i < len(s); i++ {
		if charpos[s[i]] >= start {
			start = charpos[s[i]] + 1
		} else if i+1-start > max {
			max = i+1-start
		}
		charpos[s[i]] = i
	}
	return max
}


