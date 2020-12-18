package main

import (
	"fmt"
	"net/http"

	"github.com/Go-Simple-Shop/model"
)

func main() {
	http.HandleFunc("/", model.HandleIndex)
	fmt.Println("server start at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
