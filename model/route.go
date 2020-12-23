package model

import (
	"fmt"
	"html/template"
	"mux"
	"net/http"
	"strconv"

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
func sqlQuerySelectOne(ID int) []DataBarang {
	db, err := data.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM `Daftar_Barang` WHERE ID_Barang = ?", ID)
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

func sqlQueryInput(nBarang string, nHarga int) {
	db, err := data.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("INSERT INTO `Daftar_Barang`(`ID_Barang`, `Nama_Barang`, `Harga_Barang`, `ID_User`) VALUES (?,?,?,?)", nil, nBarang, nHarga, 1)
	defer rows.Close()
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
}
func sqlDelete(ID int) {
	db, err := data.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("DELETE FROM `Daftar_Barang` WHERE `Daftar_Barang`.`ID_Barang` = ?", ID)
	defer rows.Close()
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	var Ddata = SqlQuery()
	var tmpl = template.Must(template.ParseFiles("view/index.html"))
	if err := tmpl.Execute(w, Ddata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//HandleInput Menampilkan form dari input
func HandleInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view/input.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}

func HandelProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view/input.html"))
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var nBarang = r.FormValue("nama_barang")
		var nHarga = r.FormValue("harga_barang")
		var data = map[string]string{"nama_barang": nBarang, "harga_barang": nHarga}
		cHarga, _ := strconv.Atoi(nHarga)
		sqlQueryInput(nBarang, cHarga)
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query, _ := strconv.Atoi(vars["id"])
	var Ddata = sqlQuerySelectOne(query)
	//fmt.Fprintf(w, "Category: %v\n", vars["id"])
	var tmpl = template.Must(template.ParseFiles("view/delete.html"))
	if err := tmpl.Execute(w, Ddata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func HandleDeleteProsess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query, _ := strconv.Atoi(vars["id"])
	sqlDelete(query)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
