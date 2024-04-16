package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		num1 int   = 1
		num2 int8  = 100
		num3 int16 = 10000
		num4 int32 = 1000000000
		num5 int64 = 1000000000000000000
	)

	num6 := 1000
	var num7 uint = 10

	//INT8 = Byte
	var num8 byte = 10

	//INT32 = RUNE
	var num9 rune = 10

	fmt.Println(num1, num2, num3, num4, num5, num6, num7, num8, num9)

	var num10 float32 = 10.5
	fmt.Println(num10)

	var txt1 string = "Text"
	txt2 := "Text" 
	fmt.Println(txt1, txt2)

	var bool bool
	fmt.Println(bool)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}