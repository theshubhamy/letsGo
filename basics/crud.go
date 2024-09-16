package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl string = "http://127.0.0.1:5555/items"

func crud() {
	fmt.Println("Welocme to Go Lang")
	PostReqFormData()
}

func PostReqFormData() {
	formData := url.Values{}

	formData.Add("name", "shubhams")
	formData.Add("email", "shubham@gmail.com")
	formData.Add("image", "./test.txt")
	response, err := http.PostForm(baseUrl, formData)
	checkNilErr(err)
	fmt.Println("Status Code:", response.StatusCode)
	var resString strings.Builder
	data, _ := io.ReadAll(response.Body)

	byteCount, _ := resString.Write(data)
	fmt.Println(byteCount)
	fmt.Println(resString.String())

	defer response.Body.Close()
}

func PostRequest() {
	payload := strings.NewReader((`
 {"name": "shubham kumar", "email": "shubham@example.com", "age": 24 }
 `))

	response, err := http.Post(baseUrl, "application/json", payload)
	checkNilErr(err)
	var resString strings.Builder
	data, _ := io.ReadAll(response.Body)

	byteCount, _ := resString.Write(data)
	fmt.Println(byteCount)
	fmt.Println(resString.String())

	defer response.Body.Close()
}
func GetRequest() {
	response, err := http.Get(baseUrl)
	checkNilErr(err)
	fmt.Println("Status Code:", response.StatusCode)

	var resString strings.Builder
	data, _ := io.ReadAll(response.Body)
	byteCount, _ := resString.Write(data)

	fmt.Println(byteCount)
	fmt.Println(resString.String())
	defer response.Body.Close()
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
