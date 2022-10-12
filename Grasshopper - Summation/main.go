package main

import "fmt"

func Summation(n int) (sum int) {
	//my solution was the best
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func main() {
	fmt.Println(Summation(8))
}
