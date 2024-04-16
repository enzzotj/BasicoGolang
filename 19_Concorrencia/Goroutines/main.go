package main

import (
	"fmt"
	"time"
)

func escrever(txt string){
	for{
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}

func main() {
	go escrever("oi")
	escrever("td bem?")
}