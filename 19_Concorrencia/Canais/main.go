package main

import (
	"fmt"
	"time"
)

func escrever(txt string, canal chan string){
	for i := 0; i < 5; i++{
		canal <- txt
		time.Sleep(time.Second)
	}

	close(canal)
}

func main() {
	canal := make(chan string)
	go escrever("Test", canal)

	//for{
	//mensagem, aberto := <- canal
	//if !aberto{
	//		break
	//	}
	//	fmt.Println(mensagem)
	//}

	for mensagem := range canal{
		fmt.Println(mensagem)
	}

}