package main

import (
        "fmt"
)

func main() {
        arr := []int{4,1,9,2,10,34,21,1,2, 0}
        rst := selectSort(arr)
        for _, val := range rst {
                fmt.Printf("%v ", val)
        }
}

func selectSort(input []int) []int {
        min := 0
        for i := 0; i < len(input) - 1; i++ {
                min = i
                for j := i + 1; j< len(input); j++ {
                        if input[min] > input[j] {
                                min = j
                        }
                }
                input[i], input[min] = input[min], input[i]
        }
        return input
}
