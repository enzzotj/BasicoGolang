package main

import "fmt"

type pessoa struct{
	nome string
	idade uint8
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Print("Heranca")

	p1 := pessoa{"Enzzo", 21}

	e1 := estudante{p1, "Ads", "Sptech"}

	fmt.Println(e1)
}