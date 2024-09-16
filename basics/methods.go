package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func methods() {

	fmt.Println("Structs in go lang") //no inheritance , Super,Parent

	shubham := User{"shubham", "shubham@gmail.com", 22, true}
	fmt.Println(shubham)
	shubham.GetStatus()
	shubham.NewMail()
	fmt.Println(shubham)

}

func (u User) GetStatus() {
	fmt.Println("is User Active: ", u.Status)
}

// it pass a copy of  srruct
func (u *User) NewMail() {
	u.Email = "Test@gmail.com"
	fmt.Println("User New Email is: ", u.Email)
}
