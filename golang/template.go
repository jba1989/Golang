package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/rate", loadTpl)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func loadTpl() {
	t, err := template.ParseFiles("./tpl.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
	}{
		Title: "匯率",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
