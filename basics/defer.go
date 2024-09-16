package main

import "fmt"

func defersss() {

	defer fmt.Println("Hello Defer 1")
	defer fmt.Println("Hello Defer 2")
	fmt.Println(" Defer in Go Lang")
	mydef()
}

func mydef() {
	for d := 0; d <= 5; d++ {
		defer fmt.Println(d)
	}
}
