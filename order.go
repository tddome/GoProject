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

*/
type order struct {
	dCode    string
	dPercent int
}

//Discount calculations below

func PlaceHolderWoopWoopDiscount() {
	fmt.Println("nice")
}
