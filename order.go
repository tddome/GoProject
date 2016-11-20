package main

import (
	"fmt"
	"math/rand"
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
	oUserID    int
}

var orderList map[int]int

//NewOrder - Makes new order, clears existing order
func NewOrder(uID int) {
	r := rand.New(rand.NewSource(1000))
	var uniqueID int = 0
	var count int = 0

	orderList = make(map[int]int)

	for uniqueID != 1 {
		randID = r.Int()
		for i, a := range orderHistory {
			if a.oBillID == randID {
				count++
			} else if count == len(orderHistory) {
				uniqueID = 1
			}
		}
	}

	oNew := order{
		oBillID:    randID,
		oBillTotal: 0,
		oUserID:    uID,
	}
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

func OrderQuantityUpdate(prodID int, newQ int) {
	orderList[prodID] = newQ
}
