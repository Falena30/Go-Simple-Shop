package model

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Go-Simple-Shop/data"
)

type DataBarang struct {
	ID         string
	NamaBarang string
	HargaBaang int
	IDPembuat  string
}

func SqlQuery() []DataBarang {
	db, err := data.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM Daftar_Barang")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []DataBarang

	for rows.Next() {
		var each = DataBarang{}
		var err = rows.Scan(&each.ID, &each.NamaBarang, &each.HargaBaang, &each.IDPembuat)
		if err != nil {
			fmt.Println(err.Error())
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return result
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	var Ddata = SqlQuery()

	var tmpl = template.Must(template.ParseFiles("view/index.html"))
	if err := tmpl.Execute(w, Ddata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
