package main

import (
	"fmt"
	"math"
)

// vamos criar uma função que escreva o valor das areas na tela
type forma interface { // interfaces trabalham com assinaturas de métodos
	// essa interface recebe um metodo que retorna algo
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %2.f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 { // metodo que calcula a area
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}
func main() { // na main, passamos os parametros dentro da var ao chamar a interface
	fmt.Println("Interfaces em GO")

	r := retangulo{10, 15}
	escreverArea(r) // preciso implementar o método area primeiro
	c := circulo{10}
	escreverArea(c)
}
