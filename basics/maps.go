package main

import "fmt"

// fmt.Println("maps in go lang")
// lang := make(map[int]string)

// lang[1] = "aaaaa"
// lang[2] = "bbbbb"
// lang[3] = "cccc"
// lang[4] = "dddd"
// lang[5] = "eeee"
// fmt.Println("lists of lang", lang[2])
// delete(lang, 2)
// fmt.Println("lists of lang", lang)

// // loops in goloang
//
//	for _, value := range lang {
//		fmt.Printf("for key v, value is %v\n", value)
//	}
// fmt.Println("Structs in go lang") //no inheritance , Super,Parent
// type User struct {
// 	Name   string
// 	Email  string
// 	Age    int
// 	status bool
// }
// fmt.Println("Structs in go lang") //no inheritance , Super,Parent

// 	shubham := User{"shubham", "shubham@gmail.com", 22, true}
// 	fmt.Println(shubham)

func maps() {
	fmt.Println("if else in go lang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 Login count"
	}
	fmt.Println(result)

	if num := 7; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("Num is Grater than 10")
	}
}
