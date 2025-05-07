package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("List of all language", languages)
	fmt.Println("JS shorts for", languages["js"])

	//removing from slice

	delete(languages, "rb")
	fmt.Println("List of all language after deleted", languages)

}
