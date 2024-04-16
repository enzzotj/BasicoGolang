package main

import "fmt"

func main() {
	fmt.Println(somar(10, 20));

	var f = func (txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Txt")
	fmt.Println(resultado)
	resSoma, resSub, resMult, resDiv := calculosMat(10, 15)
	fmt.Println(resSoma, resSub, resMult, resDiv)
}

func calculosMat(n1, n2 float32) (float32, float32, float32, float32){
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}