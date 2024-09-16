package main

import "fmt"

func function() {
	welcome()

	result := sum(3, 2, 34, 56, 77, 5, 000)
	fmt.Println("Sum of number", result)

}
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val

	}

	return total
}

func funChild() {
	fmt.Println("Function child")
}
func welcome() {
	fmt.Println("Function in Go Lang")
}
