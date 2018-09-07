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

func main() {
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/downloadfile", DownloadFileHandler)

	r.HandleFunc("/", IndexHandler)

	r.HandleFunc("/delete/{currency}", DeleteHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("rateLayout.html"))
	data := Flash()
	tmpl.Execute(w, data)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := Flash()
	delete(data, vars["currency"])
	fmt.Printf("%v", data)
	SaveRate(data)
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	newFileName := "rate.csv"
	file, err := os.Create(newFileName)
	defer file.Close()
	defer w.Write([]byte("finished!\n"))

	res, err := http.Get("https://rate.bot.com.tw/xrt/flcsv/0/day")
	if err != nil {
		fmt.Println("downfile error")
		return
	}
	buf := make([]byte, 1024)
	for {
		size, _ := res.Body.Read(buf)

		if size == 0 {
			break
		} else {
			file.Write(buf[:size])
		}
	}

	//data := Load("rate.csv")
	//SaveRate(data)
}

// Load download file and return a map(目前必須手動把.csv的第一行中文刪掉,不然會噴錯)
func Load(file string) map[string][]string {
	csvFile, _ := os.Open(file)
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

// Flash new file and return a map
func Flash() map[string][]string {
	csvFile, _ := os.Open("newRate.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var currency = make(map[string][]string)

	line, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	for i := 0; i < len(line); i++ {
		currency[line[i][0]] = []string{line[i][0], line[i][1], line[i][2], line[i][3], line[i][4]}
	}

	return currency
}

func SaveRate(input map[string][]string) {
	fileName := "newRate.csv"
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)

	data := input
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
