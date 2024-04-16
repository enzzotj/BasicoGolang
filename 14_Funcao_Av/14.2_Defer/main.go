package main

import "fmt"

func recupExec(){
	if r := recover(); r != nil{
		fmt.Println("Excucao recuperada com sucesso!")
	}
}

func media(n1, n2 float32) bool{
	defer recupExec()
	total := (n1 + n2) / 2

	if total > 6{
		return true
	}else{
		return false
	}
	panic("A media é exatamente 6")
}

func closure() func(){
	txt := "Dentro da função closure"

	funcao := func ()  {
		fmt.Println(txt)
	}
	return funcao
}

func inverteSinal(numero *int){
	*numero = *numero * -1
}

func init(){
	fmt.Println("Inicio de Tudo")
}

func main() {
	fmt.Println(media(8, 6))

	funcaoNova := closure()
	funcaoNova()

	numero := 10
	fmt.Println(numero)
	inverteSinal(&numero)
	fmt.Println(numero)

}