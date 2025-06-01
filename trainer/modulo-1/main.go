package main

import "fmt"

func main() {

	var idade int
	var ativo bool
	var nome string
	var altura float64

	fmt.Println(idade, ativo, nome, altura)
}

func main() {
	nome := "Carlos"
	idade := 28
	altura := 1.82
	ativo := true

	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Altura: %.2f\n", altura)

	fmt.Printf("Ativo: %b\n", ativo)
}

func main() {
	const PI float64 = 3.14

	const RAIO float64 = 5.0

	area := PI * RAIO * RAIO
	fmt.Printf("area: %.2f\n", area)

}

func main() {
	var celsius = 30.0

	F := (celsius * 1.8) + 32
	fmt.Printf("celsius: %.2f\n", F)

}
