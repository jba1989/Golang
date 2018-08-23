package main

  import (
	"os"
    "fmt"
    //"strconv"
  )

  func main() {
      //a := strconv.FormatBool(false)
      //b := strconv.FormatFloat(123.23, 'g', 12, 64)
	  //c := strconv.FormatInt(9, 8)
      //d := strconv.FormatUint(12345, 10)
	  //e := strconv.Itoa(1023)
	  f := os.Args
      fmt.Printf("%v", f)
  }