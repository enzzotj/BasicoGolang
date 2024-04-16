package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(txt string){
	for i := 0; i < 5; i++{
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func ()  {
		escrever("oi")
		waitGroup.Done()
	}()

	go func ()  {
		escrever("td bem?")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}