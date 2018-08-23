package main

import (
	"fmt"
	"os"
)

func main() {
	file := "json.go"
	f1, err := os.Open(file)
	if err != nil {
		fmt.Println(file, err)
		return
	}

	defer f1.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := f1.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}