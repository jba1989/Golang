package main

import (
	"fmt"
	"strconv"

	"./cl"
)

var a int = 1
var b int32 = 2
var c int64 = 3
var d string = "999"
var e float32 = 88.8
var f float64 = 99.9
var x string = "I LOVE Golnag_"
var sum int

func main() {
	fmt.Println("a+b=", a)
	sum = a + int(b) + int(c)
	fmt.Println("a+b+c=", sum)
	fmt.Println("f/e=", cl.Division(float64(f), float64(e)))
	fmt.Println("a+d=", a+strToInt(d))
	fmt.Println("x+a=", x+intToStr(a))
}

//數字轉文字
func intToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}
