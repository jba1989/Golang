package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Rateline struct {
	Country string  `json:"幣別"`
	Now     Now     `json:"現金"`
	Current Current `json:"即期"`
}

type Now struct {
	Buy  string `json:"本行買入"`
	Sell string `json:"本行賣出"`
}

type Current struct {
	Buy  string `json:"本行買入"`
	Sell string `json:"本行賣出"`
}

func main() {
	csvFile, _ := os.Open("rate.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var rateline []Rateline

	line, error := reader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	for i := 0; i < len(line); i++ {
		if i == 0 {
			continue
		}
		rateline = append(rateline, Rateline{
			Country: line[i][0],
			Now: Now{
				Buy:  line[i][2],
				Sell: line[i][12],
			},
			Current: Current{
				Buy:  line[i][3],
				Sell: line[i][13],
			},
		})
	}

	ratelineJson, _ := json.MarshalIndent(rateline, "", "	")
	fmt.Println(string(ratelineJson))
}
