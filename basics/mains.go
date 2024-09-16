package main

import (
	"fmt"
)

// Pointers
// func main() {
// 	fmt.Println("Hello Pointers!")

// 	myNum := 23
// 	var ptr = &myNum
// 	fmt.Println("value of pointers ", *ptr)
// 	*ptr = *ptr * 2
// 	fmt.Println(" new value of pointers ", myNum)
// }

// func main() {
// 	fmt.Println("Hello Array!")
// 	var fruitList [10]string

// 	fruitList[0] = "apple"
// 	fruitList[1] = "banna"
// 	fruitList[3] = "peach"
// 	fruitList[4] = "mango"
// 	fmt.Println("My Froui list:", len(fruitList))

// var vegList = [4]string{"a", "b", "c"}
// 	fmt.Println("My Froui list:", len(vegList))
// fmt.Printf("type of myList %T\n", vegList)

// sort
// sort.IntsAreSorted(highScores)
// fmt.Println("sorted of highScores ", sort.IntsAreSorted(highScores))

// make
// fmt.Println("type of myList ", myList)

// 	highScores := make([]int, 4)
// 	highScores[0] = 12
// 	highScores[1] = 45
// 	highScores[2] = 945
// 	highScores[3] = 95
// 	highScores = append(highScores, 11, 22, 33, 44, 55)
// 	fmt.Println("type of highScores ", highScores)
// }

func mains() {
	fmt.Println("Hello Slices!")
	var myList = []string{"a", "b", "c", "d", "e", "f"}

	fmt.Println("my list slices", myList)
	// how to remove a value from slices based on Index
	var index int = 3
	myList = append(myList[:index], myList[index+1:]...)

	fmt.Println("my list slices", myList)
}
