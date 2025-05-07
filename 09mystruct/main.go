package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Struct in golang")

	hitesh := User{"Rosada", "rosada@go.dev", true, 16}

	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
}
