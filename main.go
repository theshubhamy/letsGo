package main

import (
	"fmt"
	"net/http"

	"github.com/theshubhamy/cloverapi/routes"
)

func main() {
	fmt.Println("Let's Go")

	fmt.Println("Server is running at  4001")
	r := routes.Router()
	http.ListenAndServe(":4001", r)

}
