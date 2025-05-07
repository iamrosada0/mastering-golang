package main

import "fmt"

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// Adding items
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
}
