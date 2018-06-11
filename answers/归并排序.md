package main

import (
	"fmt"
)

func main() {
	arr :=[]int{10,5,7,1,9,2,3,6,8,4}
	fmt.Println(arr)
	rst := mergeSort(arr)
	fmt.Println(rst)
}

func mergeSort(arr []int) []int{
	length := len(arr)
	if length <=1 {
		return arr
	}

	num := length/2
	left := mergeSort(arr[:num])
	right := mergeSort(arr[num:])
	return merge(left,right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l <len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
