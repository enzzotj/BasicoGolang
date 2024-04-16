package main

import "fmt"

func main() {
	fmt.Println("If e Else")

	numero := 10

	if numero >= 15 {
		fmt.Println("E maior 15")
	}else{
		fmt.Printf("E menor 15")
	}

	if outroNum := numero; outroNum > 0{
		fmt.Println("E maior q 0");
	}
}