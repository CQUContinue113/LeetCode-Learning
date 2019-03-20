package main

import "fmt"

func main()  {
	num1 := []int{1,4}
	num2 := []int{2,3}
	fmt.Println(num1,num2)
	result := findMedianSortedArrays(num1,num2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/*for _,i := range nums1 {
		nums2 = append(nums2, i)
		fmt.Println(nums2)
	}*/
	nums2 = append(nums2,nums1...)
	fmt.Println(nums2,len(nums2))
	quicksort(nums2,0, len(nums2)-1)
	if len(nums2)%2 != 0 {
		return float64(nums2[len(nums2)/2])
	} else {
		return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2-1])/float64(2)
	}
}

func quicksort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quicksort(arr, start, j)
		}
		if end > i {
			quicksort(arr, i, end)
		}
	}
}

