package main

import "fmt"

func main() {
	str := "PAYPALISHIRING"
	fmt.Println(convert(str,3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	lens := 0
	index := 0
	num := 0
	count := numRows-1
	var arr []string
	for {
		if lens < len(s) {
			if index+numRows-1 == count {
				if lens+numRows > len(s) {
					strtemp := s[num*(2*numRows-2):len(s)]
					temp :=make([]byte,numRows)
					for t := 0; t < len(s)-num*(2*numRows-2); t++ {
						temp[t] = strtemp[t]
					}
					arr = append(arr, string(temp))
				} else {
					arr = append(arr, s[num*(2*numRows-2):num*(2*numRows-2)+numRows])
				}
				index++
				count+=numRows-1
				num++
				lens += numRows
			} else {
				pst := index%(numRows-1)
				str := make([]byte,numRows)
				str[numRows-1-pst] = s[(index/(numRows-1))*(2*numRows-2)+numRows-1+pst]
				arr = append(arr, string(str))
				index++
				lens += 1
			}
		} else {
			break
		}
	}
	var ans []byte
	for i:=0; i<numRows; i++ {
		for j:=0; j<len(arr); j++ {
			if str := arr[j]; byte(str[i]) != 0x0 {
				ans = append(ans, byte(str[i]))
			}
		}
	}

	return string(ans)
}