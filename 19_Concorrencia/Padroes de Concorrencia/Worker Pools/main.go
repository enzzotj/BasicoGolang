package main

import "fmt"

func main() {
	tarefas := make(chan int, 10)
	resultados := make(chan int, 10)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 10; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 10; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for num := range tarefas {
		resultados <- fibonacci(num)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}