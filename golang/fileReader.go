package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
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

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rateline = append(rateline, Rateline{
			Country: line[0],
			Now: Now{
				Buy:  line[2],
				Sell: line[12],
			},
			Current: Current{
				Buy:  line[3],
				Sell: line[13],
			},
		})
	}

	ratelineJson, _ := json.MarshalIndent(rateline, "", "	")
	fmt.Println(string(ratelineJson))

}
