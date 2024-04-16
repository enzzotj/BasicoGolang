package main

import "fmt"

func calcMat(n1, n2 int) (soma, sub int){
	soma =  n1 + n2
	sub = n1 - n2
	return
}

func somaTotal(numeros ...int) int{
	total := 0

	for _, num := range numeros{
		total += num
	}

	return total
}

func escrever(txt string, nums ...int){
	for _, num := range nums{
		fmt.Println(txt, num)
	}
}

func fibonacci(posicao uint) uint{
	if posicao <= 1{
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main() {
	soma, sub := calcMat(10, 20)
	fmt.Println(soma, sub)

	total := somaTotal(10, 20, 40, 73)
	fmt.Println(total)
	escrever("Oi", 5, 10)

	ret := func (txt string) string {
		return fmt.Sprintf("Recebido -> %s", txt)
	}("Ola mundo")

	fmt.Println(ret)

	fmt.Println(fibonacci(4))

}