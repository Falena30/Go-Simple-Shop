package main

import (
	"fmt"
	"mux"
	"net/http"

	"github.com/Go-Simple-Shop/model"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", model.HandleIndex)
	r.HandleFunc("/Masukkan_Barang", model.HandleInput)
	r.HandleFunc("/process", model.HandelProcess)
	r.HandleFunc("/delete/{id}", model.HandleDelete)
	r.HandleFunc("/delete/prosess/{id}", model.HandleDeleteProsess)
	fmt.Println("server start at localhost:8080")
	http.ListenAndServe(":8080", r)
}
