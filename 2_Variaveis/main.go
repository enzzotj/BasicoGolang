package main

import "fmt"

func main() {
	var variavel1 string = "Var1"
	variavel2 := "Var2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "Var3"
		variavel4 string = "Var4"
	)

	fmt.Println(variavel3, variavel4);

	variavel5, variavel6 := "Var5", "Var6"

	fmt.Println(variavel5, variavel6);

	const constante1 string = "Constante1"

	fmt.Println(constante1)
}