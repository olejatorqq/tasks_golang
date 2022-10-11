package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) (powers []uint64) {
	//best solution
	for i := 0; i <= n; i++ {
		powers = append(powers, uint64(math.Pow(2, float64(i))))
	}
	return

	//my solution

	//var resultArray []uint64
	//for i := 0; i < n+1; i++ {
	//	resultArray = append(resultArray, uint64(math.Pow(2, float64(i))))
	//}
	//return resultArray
}

func main() {
	fmt.Println(PowersOfTwo(0))
}
