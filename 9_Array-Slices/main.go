package main

import "fmt"

func main() {
	fmt.Println("Array e Slices")

	var ar1 [5]int
	ar1[0] = 1
	fmt.Println(ar1)

	ar2 := [5]string{"a", "b"}
	fmt.Println(ar2)

	slice := []int{1, 2}
	slice = append(slice, 1)
	fmt.Println(slice)

	slice2 := ar2[1:3]
	fmt.Println(slice2)

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
}