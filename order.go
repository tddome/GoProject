package main

import (
	"fmt"
)

/*
Order - Stores a user/manager Order

Generates:
Bill that can be accessed by User and Supplier
Prints Bill information which contains:
	Bill ID
	Customer ID
	Bank Account (Last 4 digits)
	Products purchased with quantities for each
	Applied discount(s)
	Total Cost

What comes from this class:
	Map, with key being Product and value being Quantity
		List of product IDs user wants in the order (can add/remove)
		Number of products user wants for the item
	Bill ID to identify this bill (Basically order number)
	Bill Total to state total price (Can be modified by discount)

What should come from other classes:
	Customer ID
	Bank Account
	Applied Discounts (Discount class modifies total)
*/

type order struct {
	oBillID    int
	oBillTotal int
}

var orderList map[int]int

//NewOrder - Makes new order, clears existing order
func NewOrder() {
	orderList = make(map[int]int)
}

func PrintOrder() {
	for opID, oQuantity := range orderList {
		fmt.Println("Product ID: ", opID, "Quantity: ", oQuantity)
	}
}

func OrderAddProduct(prodID int, prodQ int) {
	orderList[prodID] = prodQ
}

func OrderRemoveProduct(prodID int, prodQ int) {
	delete(orderList, prodID)
}
