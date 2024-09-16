package main

import (
	"fmt"
	"io"
	"os"
)

func files() {
	fmt.Println("Files in go Lang")

	content := "This need to go in files - thehsubham.dev"

	files, err := os.Create("./myfiles.txt")
	checkNilErr(err)
	length, err := io.WriteString(files, content)
	checkNilErr(err)

	fmt.Println("length is :", length)
	readFile("./myfiles.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("file data", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
