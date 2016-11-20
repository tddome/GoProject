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
	List of products user wants in the order (can add/remove)
	Number of products user wants for the item
		(INDEXES MUST MATCH IN BOTH OF THESE ARRAYS!)
		(e.g. oProductList[0] = "Steak", oQuantity[0] = 3)
		(Buying 3 steaks)
	Bill ID to identify this bill (Basically order number)
	Bill Total to state total price (Can be modified by discount)

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
