package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Falsh() {
	tmpl := template.Must(template.ParseFiles("rateLayout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Load()
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/de/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars["id"])
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

// Load will return a map
func Load() map[string][]string {
	csvFile, _ := os.Open("rate.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var currency = make(map[string][]string)

	line, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	for i := 0; i < len(line); i++ {
		if i == 0 {
			continue
		}
		currency[line[i][0]] = []string{line[i][0], line[i][2], line[i][12], line[i][3], line[i][13]}
	}

	return currency
}

// Flash will return a map
func Flash() map[string][]string {
	csvFile, _ := os.Open("newRate.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var currency = make(map[string][]string)

	line, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	for i := 0; i < len(line); i++ {
		if i == 0 {
			continue
		}
		currency[line[i][0]] = []string{line[i][0], line[i][1], line[i][2], line[i][3], line[i][4]}
	}

	return currency
}

/*
func DeleteCurrency() {
	_, ok := Currency["CAD"]
	if ok != nil {
		delete(Currency, "CAD")
	}

}
*/

func SaveRate() {
	fileName := "newRate.csv"
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)

	data := Load()
	for _, v := range data {
		r2.Write(v)
		r2.Flush()
	}

	fout, err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	fout.WriteString(buf.String())
}
