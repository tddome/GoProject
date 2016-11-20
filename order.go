package main

import (
	"fmt"
)

/*
Order - Stores a user/manager Order
Contains:
Ability to add products (orderproducts product)
Ability to add product quantity (quantity

Generates:
Bill that can be accessed by User and Supplier
Prints Bill information which contains:
	Bill ID
	Customer ID
	Bank Account (Last 4 digits)
	Products purchased with quantities for each
	Applied discount(s)
	Total Cost

What should come from other classes:
	Customer ID
	Bank Account
	Applied Discounts (Discount class modifies total)

*/
type order struct {
	oProductList []product
	oQuantity    []int
	oBillID      int
	oBillTotal   int
}

//Discount calculations below

func PlaceHolderWoopWoopOrder() {
	fmt.Println("nice")
}
