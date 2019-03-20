package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	str := "-91283472332"
	fmt.Println(myAtoi(str))
}

func myAtoi(str string) int {
	length := len(str)
	var temp1 float64
	var ans []int
	var ret float64
	temp := false
	flag1 := false
	cont := true
	for i := 0; i < length; i++ {
		if (str[i] == 32 && flag1 == false)||cont == false{
			if cont == false {
				break
			}
			continue
		} else {
			if (str[i] < 48 || str[i] > 57) && str[i]!=45 &&str[i]!=43{
				return 0
			} else if str[i] == 45{
				temp = true
				cont = false
				if i + 1 >= length {
					break
				} else if str[i+1] == 43{
					temp = false
					cont = false
				} else if str[i+1] >= 48 && str[i+1] <=57 {
					cont = true
				} else if str[i+1] < 48 || str[i+1] >57{
					cont = false
				}
			}else if str[i] == 43 {
				temp = false
				cont = false
				if i + 1 >= length {
					break
				} else if str[i+1] == 45{
					temp = true
					cont = false
				} else if str[i+1] >= 48 && str[i+1] <=57 {
					cont = true
				} else if str[i+1] < 48 || str[i+1] >57{
					cont = false
				}
			} else if cont == true {
				t, _ := strconv.Atoi(string(str[i]))
				ans = append(ans, t)
				if i + 1 >= length {
					break
				} else if str[i+1] < 48 || str[i+1] > 57 {
					cont = false
				}
			} else {
				break
			}
		}
	}
	for i := 0; i < len(ans); i++ {
		temp1 = float64(math.Pow10(len(ans)-1-i))
		ret = ret + float64(temp1)*float64(ans[i])
	}
	/*if temp == true {
		ret = ret*(-1)
	}*/
	if temp==false&&ret > float64(math.Pow(2,31)-1) {
		return int(math.Pow(2,31)-1)
	} else if temp == true&&ret*(-1)<float64(math.Pow(2,31)*(-1)) {
		return int(math.Pow(2,31)*(-1))
	} else {
		if temp == true {
			return int(ret*(-1))
		}
		return int(ret)
	}
}
