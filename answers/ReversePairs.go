package main

import (
	"fmt"
)

func main() {
	arr := []int{2,4,1,3,5}
	fmt.Println("count:", reversePairs(arr))
	return
}

func reversePairs(arr []int) int {
	var i, j int
	count := 0
	for i=0; i< len(arr) - 1; i ++ {
		for j = i + 1 ; j < len(arr) ; j ++ {
			if  arr[j] < arr[i]  {
				fmt.Printf("%d  ", arr[i])
				fmt.Printf("%d\n", arr[j])
				count ++ 
			}
		}
	}
	return count
}
