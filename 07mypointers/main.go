package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointers class")

	// var ptr *int

	// fmt.Println("Value of pointers is", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of actual pointers is", ptr)

	fmt.Println("Value of actual pointers is", *ptr)

	*ptr = *ptr * 2

	fmt.Println("The new value is", myNumber)

}
