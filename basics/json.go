package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Author   string   `json:"author"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func mainjson() {

	fmt.Println("JSON Playgroiund in Go Lang")
	DecodeJson()
}

func EncodeJson() {
	newCourse := []course{{"Go lang", 299, "shubhma", "abc", []string{"dev", "js"}}, {" react", 2999, "shubhma", "abc", []string{"dev", "js"}}, {"Node", 599, "shubhma", "abc", []string{"dev", "js"}}, {"React Native ", 2299, "shubhma", "abc", []string{"dev", "js"}}}

	finalJson, err := json.MarshalIndent(newCourse, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonData := []byte(`
 				{
                "coursename": "Go lang",
                "price": 299,
                "author": "shubhma",
                "tags": [
                        "dev",
                        "js"
                ]
        }
`)

	var newCourses course
	isJson := json.Valid(jsonData)
	if !isJson {
		fmt.Println("not valid josn format")
	}

	json.Unmarshal(jsonData, &newCourses)
	fmt.Printf("%#v\n", newCourses)

	var onlineData map[string]interface{}
	json.Unmarshal(jsonData, &onlineData)
	// fmt.Printf("%#v\n", onlineData)

	for key, val := range onlineData {
		fmt.Printf("Key is : %v , value is: %v, datatype is: %T\n", key, val, val)
	}
}
