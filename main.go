package main

import (
	"fmt"
	"os"
)

func main() {

	/*
	***********************************************
	* This is the driver code. Don't change it!!!
	***********************************************
	 */
	orders := []string{
		"Regular Electronics:100:2 Books:20:1",
		"Member Clothing:50:4",
		"VIP Electronics:100:1 Clothing:50:2",
		"Gold Books:20:1",
		"Regular Toys:30:2",
	}

	if len(os.Args) == 0 {
		fmt.Println("Arguments are required to run the program.")
		return
	}
	input := os.Args[1]
	for _, order := range orders {
		if order == input {
			Handle(order)
			return
		}
	}

}
