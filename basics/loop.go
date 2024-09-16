package main

import "fmt"

func loops() {
	fmt.Println("Loops In Go Lang")

	days := []string{"sun", "Tue", "Fri", "Sat"}
	fmt.Println(days)

	//for loop

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println((days[d]))
	// }

	// for rnage loop
	// for d := range days {
	// 	fmt.Println((days[d]))
	// }

	// for index, day := range days {
	// 	fmt.Println("index : ", index, "day is : ", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco //jump to lco
		}
		if rougueValue == 5 {
			rougueValue++
			break
		}
		fmt.Println("vaue is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("value is 2")

}
