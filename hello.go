package main

import "fmt"

func main() {
	var userUrl string

	fmt.Println("Enter a fucking url")
	fmt.Scan(&userUrl)

	if isUrlReal(userUrl) {
		fmt.Println("Such a good boy")
	} else {
		fmt.Println("IT'S NOT A FUCKING URL")
	}
}
