package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// Adding items
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 232
	highScores[1] = 231
	highScores[2] = 233
	highScores[3] = 234
	//highScores[4] = 934
	highScores = append(highScores, 555, 566, 455)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

}
