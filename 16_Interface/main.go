package main

import (
	"fmt"
	"math"
)

type retangulo struct{
	altura float64
	largura float64
}

type circulo struct{
	raio float64
}

type forma interface{
	area() float64
}

func escreverArea(f forma){
	fmt.Println("A area é %0.2f", f.area())
}

func (r retangulo) area() float64{
	return r.altura * r.largura
}

func (c circulo) area() float64{
	return math.Pi * (c.raio * c.raio)
}

func main() {
	fmt.Println()
	r := retangulo{10, 15}
	escreverArea(r)
	c := circulo{20}
	escreverArea(c)
}