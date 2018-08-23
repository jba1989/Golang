package main

import "fmt"

type Employee struct {
	ID    int
	name  string
	phone string
	addr  string
}

var info Employee

func main() {
	info.ID = 1001
	info.name = "cl"

	fmt.Println(info.ID, info.name)

	fmt.Println(sum(ID *info, ID *info))
}

func sum(x, y int) int {
	z := x * 10 + y * 100
	return z
}