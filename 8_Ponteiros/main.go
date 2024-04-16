package main

import "fmt"

func main() {
	fmt.Println("Ponteiro")

	var v1 int = 10
	var v2 int = v1

	v1++
	fmt.Println(v1, v2)
	var v3 int = 15
	var pont *int

	pont = &v3
	v3++
	fmt.Println(v3, *pont)
}