package main

import (
	"bufio"
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Rateline struct {
	Country     string
	NowBuy      string
	NowSell     string
	CurrentBuy  string
	CurrentSell string
}

func main() {

	csvFile, _ := os.Open("rate.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var currency []Rateline

	line, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	for i := 0; i < len(line); i++ {
		if i == 0 {
			continue
		}
		currency = append(currency, Rateline{
			Country:     line[i][0],
			NowBuy:      line[i][2],
			NowSell:     line[i][12],
			CurrentBuy:  line[i][3],
			CurrentSell: line[i][13],
		})
	}

	tmpl := template.Must(template.ParseFiles("rateLayout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := currency
		tmpl.Execute(w, data)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
