package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Handle(input string) {
	parts := strings.Split(input, " ")

	customer := parts[0]

	discount := 0.0
	if customer == "Regular" {
		discount = 0
	} else if customer == "Member" {
		discount = 10
	} else if customer == "VIP" {
		discount = 20
	} else {
		fmt.Println("INVALID CUSTOMER TYPE")
		return
	}

	subtotal := 0.0
	shipping := 0

	for i := 1; i < len(parts); i++ {
		item := strings.Split(parts[i], ":")

		category := item[0]
		price, _ := strconv.Atoi(item[1])
		quantity, _ := strconv.Atoi(item[2])

		subtotal += float64(price * quantity)

		if category == "Electronics" {
			shipping += 10 * quantity
		} else if category == "Books" {
			shipping += 0
		} else if category == "Clothing" {
			shipping += 5 * quantity
		} else {
			fmt.Println("INVALID CATEGORY")
			return
		}
	}

	total := subtotal - (subtotal * discount / 100)

	if customer != "VIP" {
		total += float64(shipping)
	}

	fmt.Printf("Order Total: %.2f\n", total)
}
