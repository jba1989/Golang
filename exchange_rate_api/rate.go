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
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// 匯率總表
	r.HandleFunc("/", IndexHandler)

	// 下載匯率檔案
	r.HandleFunc("/download", DownloadFileHandler)

	// 選擇單一幣別
	r.HandleFunc("/pick/{currency}", PickHandler)

	// 增加一幣別匯率 nb:現金買, ns:現金賣, cb:即期買, cs:即期賣
	r.HandleFunc("/add/{currency}/{nb}/{ns}/{cb}/{cs}", AddHandler)

	// 刪除單一幣別
	r.HandleFunc("/delete/{currency}", DeleteHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}

// 匯率總表
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("rateLayout.html", "css.html"))
	data := Flash()
	tmpl.Execute(w, data)
}

// 下載匯率檔案(目前必須手動把.csv的第一行中文刪掉,不然會噴錯)
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	newFileName := "rawRate.csv"
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

	removeContentTitle()
	data := Load()
	SaveRate(data)
}

// 選擇單一幣別
func PickHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := Flash()
	rate, ok := data[vars["currency"]]
	if ok != true {
		defer w.Write([]byte("No this currency\n"))
	}
	tmpl := template.Must(template.ParseFiles("singleRateLayout.html", "css.html"))
	tmpl.Execute(w, rate)
}

// 增加一幣別匯率
func AddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := Flash()
	data[vars["currency"]] = []string{vars["currency"], vars["nb"], vars["ns"], vars["cb"], vars["cs"]}
	SaveRate(data)
	tmpl := template.Must(template.ParseFiles("rateLayout.html", "css.html"))
	tmpl.Execute(w, data)
}

// 刪除單一幣別
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := Flash()
	_, ok := data[vars["currency"]]
	if ok != true {
		defer w.Write([]byte("No this currency\n"))
	}
	delete(data, vars["currency"])
	SaveRate(data)
	tmpl := template.Must(template.ParseFiles("rateLayout.html", "css.html"))
	tmpl.Execute(w, data)
}

// 將csv的第一行中文刪掉(不然在csv套件讀取時會噴錯)
func removeContentTitle() {
	fileName := "rawRate.csv"
	newFileName := "rate.csv"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	newFile, err := os.Create(newFileName)
	defer file.Close()
	defer newFile.Close()

	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}

	buf := bufio.NewReader(file)
	i := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return
		}

		if i == 0 {
			i++
			continue
		}
		line = strings.TrimSpace(line)
		newFile.WriteString(line + "\n")
	}
}

// 讀rate.csv的檔案並回傳各幣別的匯率
func Load() map[string][]string {
	csvFile, _ := os.Open("rate.csv")
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var currency = make(map[string][]string)

	line, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	for i := 0; i < len(line); i++ {
		currency[line[i][0]] = []string{line[i][0], line[i][2], line[i][12], line[i][3], line[i][13]}
	}

	return currency
}

// 讀取已newRate.csv並回傳各幣別的匯率
func Flash() map[string][]string {
	csvFile, _ := os.Open("newRate.csv")
	defer csvFile.Close()
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

// 將匯率寫進newRate.csv內
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
