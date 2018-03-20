package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 11; i++ {
		fmt.Printf("Fibonacci(%d):%d", i, Fibonacci(i))
		fmt.Printf("\n")
	}
	return
}


func Fibonacci(n int) int32 {
	if n <= 0 {
		return -1 
	} else if n ==1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
