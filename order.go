/* GoProject
*
* Class:
*	CS 408
*
* Contributors:
* 	Troy Dome
* 	Shaylyn Wetts
*
* Last Updated:
*	11/20/2016
*
* Function:
*	Creates an order.  Creates a map of all inputted
*	products and quantities, and stores the final bill ID,
*	total price, and user ID who created the order in an
*	order struct.  This struct is then stored in the database.
 */

package main

import (
	"fmt"
	"math/rand"
)

// Order - stores ID, total price, and user ID
type order struct {
	oBillID    int
	oBillTotal float32
	oUserID    int
}

// Map of current ordered products and quantities
var orderList map[int]int

// GetIndexOfOrder - Check if an order is contained by ID
// and returns the index where the ID is found, else return -1
func GetIndexOfOrder(id int) int {
	for i, a := range orderHistory {
		if a.oBillID == id {
			return i
		}
	}
	return -1
}

// NewOrder - Makes new order, clears existing order
func NewOrder() {
	orderList = make(map[int]int)
}

// PrintOrder - Prints all current products and quantities in
// the order map
func PrintOrder() {
	for opID, oQuantity := range orderList {
		fmt.Println("Product ID: ", opID, "; Quantity: ", oQuantity)
		fmt.Println("\n")
	}
}

// OrderAddProduct - Adds a new product and quantity to the current
// order map
func OrderAddProduct(prodID int, prodQ int) {
	orderList[prodID] = prodQ
}

// OrderRemoveProduct - Removes a product from the current order
// map
func OrderRemoveProduct(prodID int) {
	delete(orderList, prodID)
}

// OrderQuantityUpdate - Updates the quantity of a product currently
// in the map
func OrderQuantityUpdate(prodID int, newQ int) {
	orderList[prodID] = newQ
}

// CalculateOrder - Returns a value for the total of all current
// products in the map
func CalculateOrder() float32 {
	var totalOfOrder = 0.0

	for opID, oQuantity := range orderList {
		fmt.Println("Product ID: ", opID, "Quantity: ", oQuantity)
		totalOfOrder = float64(totalOfOrder) +
			float64(GetProductPrice(opID)*float32(oQuantity))
	}
	return float32(totalOfOrder)
}

// FinalizeOrder - Finalizes the current order; calls CalculateOrder
// to get the total of the current products, applies discounts based
// on inputted discount code, and generates an order object to be
// stored in the database
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

	bTotal = bTotal - DiscountAmount(bTotal, discountCode)

	oNew := order{
		oBillID:    randID,
		oBillTotal: bTotal,
		oUserID:    uID,
	}

	AddOrderToDatabase(oNew)
	OrderToString(oNew.oBillID)
}

// OrderToString - Prints the finalized order information
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
