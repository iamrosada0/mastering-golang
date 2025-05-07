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

	//Removing the value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python"}

	fmt.Println(courses)

	var index int = 2
	//In this section, the slice courses initially contains 4 elements. The goal is to remove the element at index 2 ("swift"). By using the append function to concatenate the slice before and after the element to be removed, it effectively removes "swift" from the courses slice.
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
