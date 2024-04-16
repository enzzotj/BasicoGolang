package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10{
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	for j:= 0; j < 10; j++{
		fmt.Println(j)
	}

	nomes := [3]string{"Enzzo", "Jaco", "TT"}

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	for i, letra := range "Test"{
		fmt.Println(i, string(letra))
	}

	user := map[string]string{
		"nome" : "Enzzo",
		"sobrenome": "Jaco",
	}

	for chave, value := range user{
		fmt.Println(chave, value)
	}
}