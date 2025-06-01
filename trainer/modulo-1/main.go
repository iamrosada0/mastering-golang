package main

import "fmt"

func exercicio1() {
	var idade int
	var ativo bool
	var nome string
	var altura float64

	fmt.Println("Exercício 1:")
	fmt.Println("Idade:", idade)   // 0
	fmt.Println("Ativo:", ativo)   // false
	fmt.Println("Nome:", nome)     // ""
	fmt.Println("Altura:", altura) // 0.0
	fmt.Println()
}

func exercicio2() {
	nome := "Carlos"
	idade := 28
	altura := 1.82
	ativo := true

	fmt.Println("Exercício 2:")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Altura: %.2f\n", altura)
	fmt.Printf("Ativo: %t\n", ativo)
	fmt.Println()
}

func exercicio3() {
	const PI float64 = 3.14
	const RAIO float64 = 5.0

	area := PI * RAIO * RAIO

	fmt.Println("Exercício 3:")
	fmt.Printf("Área do círculo: %.2f\n", area)
	fmt.Println()
}

func exercicio4() {
	var celsius = 30.0
	fahrenheit := (celsius * 1.8) + 32

	fmt.Println("Exercício 4:")
	fmt.Printf("%.2f °C = %.2f °F\n", celsius, fahrenheit)
	fmt.Println()
}

func main() {
	exercicio1()
	exercicio2()
	exercicio3()
	exercicio4()
}
