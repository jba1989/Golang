package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	newFileName := "rate.csv"
	file, err := os.Create(newFileName)
	defer file.Close()

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

}
