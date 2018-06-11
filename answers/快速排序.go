package main

import(
	"fmt"
)

func main() {
	arr := []int{10,1,4,6,5,2,7,3,8,9}
	fmt.Println("begin", arr)
	quickSort(arr, 0, len(arr) -1)
	fmt.Println("end",arr)
}

func quickSort(arr []int, left, right int) {
	if left > right {
		return
	}

	i, j := left, right
	temp := arr[left]
	for ; i!=j; {
		for ; arr[j] >= temp && i < j; {
			j--
		}
		for ; arr[i] <= temp && i < j; {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left] = arr[i]
	arr[i] = temp
	quickSort(arr, left, i - 1 )
	quickSort(arr, i + 1, right)
}
