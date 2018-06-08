package main

import (
  "fmt"
)
func main() {
    inputArr := []int{4,1,9,2,10,34,21}
    SelectionSort(inputArr)
    return
}

func SelectionSort(nums []int) {
    length := len(nums)
    for i := 0; i < length; i++ {
        min := i
        for j := i + 1; j < length; j++ {
            if nums[j] < nums[min] {
                min = j
            }
        }
        temp := nums[i]
        nums[i] = nums[min]
        nums[min] = temp
    }
    fmt.Println(nums)
}
