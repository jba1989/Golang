package main

import "strconv"

func strToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
