package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const myurl string = "https://theshubham.hashnode.dev/displaying-pdf-in-react-js-using-adobe-pdf-embed-api"

func webURL() {
	fmt.Println("Web Request in Go lang")

	resultUrl, _ := url.Parse(myurl)
	fmt.Println(resultUrl)

	fmt.Println(resultUrl.Scheme)
	fmt.Println(resultUrl.Host)
	fmt.Println(resultUrl.Path)
	fmt.Println(resultUrl.Port())
	fmt.Println(resultUrl.RawQuery)
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("response type: %T\n", response)

	// data, errs := io.ReadAll(response.Body)
	// if errs != nil {
	// 	panic(errs)
	// }
	// content := string(data)
	// fmt.Println(content)

}
