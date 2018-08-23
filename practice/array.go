package main

import "fmt"

type Currnet int

var m = make(map[string]int)

const (
	USD Currnet = iota
	EUR
	GBP
)

func main() {
	symbol := [...]string{USD: "us", EUR: "eu", GBP: "gb"}
	for k, v := range symbol {
		fmt.Println(k, v)
	}
}



func k(list []string) string { return fmt.Sprintf("%q", list) }