package main

import (
  "fmt"
)
func main() {
    inputArr := []int{4,1,9,2,10,34,21,1,2, 0}
    BubbleSort(inputArr)
    return
}

func BubbleSort(values []int) {
	for i := 0; i < len(values) -1; i++ {
		flag := true
		for j := 0; j <  len(values) -1 -i; j ++ {
			if values[j] > values[j+1] {
				flag = false
				values[j], values[j + 1] = values[j+1],values[j]
			}
		}
		if flag == true {
			break
		}
	}
	fmt.Println(values)
}
