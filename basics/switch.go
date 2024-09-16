package main

import (
	"fmt"
	"math/rand"
	"time"
)

func schs() {

	fmt.Println("Switch and case in Go Lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, you can open")
	case 2:
		fmt.Println("Dice value is 2, you can open")
	case 3:
		fmt.Println("Dice value is 3, you can open")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4, you can open")
	case 5:
		fmt.Println("Dice value is 5, you can open")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6, you can open")
	default:
		fmt.Println("Dice value is default, you can open")
	}
}
