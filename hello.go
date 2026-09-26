package main

import (
	"fmt"
	"log"
)

func main() {
	var userUrl string

	fmt.Println("Enter a fucking url")
	fmt.Scan(&userUrl)

	res, err := isUrlReal(userUrl)

	if err != nil {
		log.Fatal("FATAL")
	}

	if res {
		fmt.Println("Such a good boy")
	} else {
		fmt.Println("IT'S NOT A FUCKING URL")
	}
}
