package main

import "fmt"

type usuarios struct{
	nome string
	idade uint8
	endereco endereco
}

type endereco struct{
	rua string
	numero string
}

func main() {
	fmt.Println("Arquivos Structs")

	var u usuarios
	u.nome = "Enzzo"
	u.idade = 21
	fmt.Println(u)

	u2 := usuarios{"Test", 1, endereco{"test", "12A"}}
	fmt.Println(u2)

	u3 := usuarios{idade: 5}
	fmt.Println(u3)
}