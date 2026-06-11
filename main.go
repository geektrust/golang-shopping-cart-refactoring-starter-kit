package main

import "fmt"

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
	args := []string{} // Replace with os.Args[1:] in actual usage
	if len(args) == 0 {
		fmt.Println("Arguments are required to run the program.")
		return
	}
	input := args[0]
	for _, order := range orders {
		if order == input {
			Handle(order)
			return
		}
	}

}
