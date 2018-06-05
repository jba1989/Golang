package main

import "strconv"

//文字轉數字
func strToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
