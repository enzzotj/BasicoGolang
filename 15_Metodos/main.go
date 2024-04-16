package main

import "fmt"

type usuarios struct{
	nome string
	idade uint8
}

func (u usuarios) salvar(){
	fmt.Println("Salvando")
}

func main() {
	u := usuarios{"Enzzo", 21}
	u.salvar()
}