package main

import (
	"fmt"
	"reflect"
)

var x float64 = 3.4

func main() {
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind())
	fmt.Println("value:", v.Float())
}
