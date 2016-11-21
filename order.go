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
	oBillTotal float32
	oUserID    int
}

var orderList map[int]int

func GetIndexOfOrder(id int) int {
	for i, a := range orderHistory {
		if a.oBillID == id {
			return i
		}
	}
	return -1
}

//NewOrder - Makes new order, clears existing order
func NewOrder() {
	orderList = make(map[int]int)
}

func PrintOrder() {
	for opID, oQuantity := range orderList {
		fmt.Println("Product ID: ", opID, "; Quantity: ", oQuantity)
		fmt.Println("\n")
	}
}

func OrderAddProduct(prodID int, prodQ int) {
	orderList[prodID] = prodQ
}

func OrderRemoveProduct(prodID int) {
	delete(orderList, prodID)
}

func OrderQuantityUpdate(prodID int, newQ int) {
	orderList[prodID] = newQ
}

func CalculateOrder() float32 {
	var totalOfOrder = 0.0
	//Calculate total amount that's in the map
	for opID, oQuantity := range orderList {
		fmt.Println("Product ID: ", opID, "Quantity: ", oQuantity)
		totalOfOrder = float64(totalOfOrder) + float64(GetProductPrice(opID)*float32(oQuantity))
	}
	//Return a total amount
	return float32(totalOfOrder)
}

func FinalizeOrder(uID int, discountCode string) {
	var count int = 0
	var randID int = 0
	var bTotal float32 = CalculateOrder()

	randID = rand.Intn(1000)
	for count < len(orderHistory) {
		if orderHistory[count].oBillID == randID {
			randID = rand.Intn(1000)
			count = 0
		} else {
			count++
		}
	}

	// function call here for generating bill total and setting bTotal value
	bTotal = bTotal - DiscountAmount(bTotal, discountCode)

	oNew := order{
		oBillID:    randID,
		oBillTotal: bTotal,
		oUserID:    uID,
	}

	AddOrderToDatabase(oNew)
	OrderToString(oNew.oBillID)
}

func OrderToString(id int) {
	var count = 0
	var i int
	i = GetIndexOfOrder(id)

	fmt.Println("Order # ", count)
	fmt.Println("Bill ID:                    ", orderHistory[i].oBillID)
	fmt.Println("Bill Total:                ", orderHistory[i].oBillTotal)
	fmt.Println("Created by User ID: ", orderHistory[i].oUserID)
	fmt.Println("\n")
}
