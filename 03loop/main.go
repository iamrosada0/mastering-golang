package main

import "fmt"

func main() {

	for i := 0; i < 1000; i++ {

		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")

	}
	likeWhile()
}

func likeWhile() {
	fmt.Println()

	i := 10

	for {
		if i < 0 {
			break
		}

		fmt.Print(i, " ")
	}
}
