package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func main() {
	if i := f(5); i == 0 {
		fmt.Println("i=", 100)
	}
}

func f(i int) int {
	i *= 0
	return i
}
