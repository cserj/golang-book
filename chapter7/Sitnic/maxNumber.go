package main

import "fmt"

func A(x ...int) int {
	maxNumber := x[0]
	for _, y := range x {
		if y > maxNumber {
			maxNumber = y
		}
	}

	return maxNumber
}

func main() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(A(x...))
}
