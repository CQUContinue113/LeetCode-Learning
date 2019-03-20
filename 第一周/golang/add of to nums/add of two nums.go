package main

import (
	"fmt"
)

type LinkList struct {
	Val int
	Next *LinkList
}

func main()  {
	num1 := LinkList{8,nil}
	num2 := LinkList{1,&num1}
	//num3 := LinkList{2,&num2}

	num4 := LinkList{0,nil}
	//num5 := LinkList{6,&num4}
	//num6 := LinkList{5,&num5}

	res := addTwoNumbers(&num2,&num4)
	fmt.Println(res)
}

func addTwoNumbers(l1 *LinkList, l2 *LinkList) *LinkList {
	var p0,p1 LinkList
	ans := &LinkList{0,nil}
	carry := 0
	p0.Next = l1
	p1.Next = l2
	curr := ans
	for p0.Next!=nil||p1.Next!=nil {
		var x,y int
		if p0.Next!=nil&&p1.Next!=nil {
			x = p0.Next.Val
			y = p1.Next.Val
		}  else if p0.Next == nil {
			x = 0
			y = p1.Next.Val
		} else if p1.Next == nil {
			y = 0
			x = p0.Next.Val
		}
		curr.Next = &LinkList{(x+y+carry)%10,nil}
		curr = curr.Next
		if x+y+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		if p1.Next.Next == nil&&p0.Next.Next==nil {
			p1.Next = nil
			p0.Next = nil
		} else if p0.Next.Next == nil {
			p0.Next = &LinkList{0,nil}
			p1.Next = p1.Next.Next
		} else if p1.Next.Next == nil {
			p0.Next = p0.Next.Next
			p1.Next = &LinkList{0,nil}
		} else {
			p1.Next = p1.Next.Next
			p0.Next = p0.Next.Next
		}
	}
	if carry==1 {
		curr.Next = &LinkList{carry,nil}
	}
	return ans.Next
}