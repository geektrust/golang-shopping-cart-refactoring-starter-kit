package main

import (
	"fmt"
	"strconv"
	"strings"
)

var discounts = map[string]float64{
	"Regular": 0,
	"Member":  10,
	"VIP":     20,
}

var shippingCosts = map[string]int{
	"Electronics": 10,
	"Books":       0,
	"Clothing":    5,
}

func Handle(input string) {
	parts := strings.Split(input, " ")

	customer := parts[0]

	discount, ok := discounts[customer]
	if !ok {
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

		shippingCost, ok := shippingCosts[category]
		if !ok {
			fmt.Println("INVALID CATEGORY")
			return
		}

		subtotal += float64(price * quantity)
		shipping += shippingCost * quantity
	}

	total := subtotal - (subtotal * discount / 100)

	if customer != "VIP" {
		total += float64(shipping)
	}

	fmt.Printf("Order Total: %.2f\n", total)
}
