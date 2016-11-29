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
*	11/28/2016
*
* Function:
*	Creates a manager.  Includes functions used by the
*	manager for adding products, removing products, adding
*	discounts, removing discounts, and accessing all current
*	products, discounts, and orders.  Includes prompts for
*	user input to run these functions.
 */

package main

import (
	"fmt"
)

// Manager - stores ID and password
type manager struct {
	mID string
	mPW string
}

// CheckProductList - Outputs information on all current
// products
func CheckProductList() {
	for _, prod := range productList {
		ProductToString(prod.pID)
	}
}

// CheckDiscountList - Outputs  information on all
// current discounts
func CheckDiscountList() {
	for _, dis := range discountList {
		DiscountToString(dis.dCode)
	}
}

// CheckOrderList - Outputs information on all currently
// stored orders
func CheckOrderList() {
	for _, ord := range orderHistory {
		OrderToString(ord.oBillID)
	}
}

// ManagerAuthenticated - All functionality for a manager
func ManagerAuthenticated() {
	fmt.Println("\nWelcome, manager!\n")
	var mAuth = 0

	for mAuth < 1 {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Add a product")
		fmt.Println("2. Delete a product")
		fmt.Println("3. Access product list")
		fmt.Println("4. Add a discount code")
		fmt.Println("5. Remove a discount code")
		fmt.Println("6. Access discount list")
		fmt.Println("7. Access order list")
		fmt.Println("8. Sign out\n")
		var text string
		fmt.Printf("-> ")
		fmt.Scan(&text)

		switch text {
		case "1":
			var inName string
			var inId int
			var inType string
			var inPrice float32

			fmt.Printf("\nEnter Name: ")
			fmt.Scan(&inName)
			fmt.Printf("Enter ID: ")
			fmt.Scan(&inId)
			fmt.Printf("Enter Type: ")
			fmt.Scan(&inType)
			fmt.Printf("Enter Price: ")
			fmt.Scan(&inPrice)
			fmt.Println("\nCreating product...")
			CreateProduct(inName, inId, inType, inPrice)
			fmt.Println("Product created.\n")
		case "2":
			var inId int

			fmt.Printf("\nEnter Product ID: ")
			fmt.Scan(&inId)
			fmt.Println("Removing product...\n")
			DeleteProduct(inId)
			fmt.Println("Product removed.\n")
		case "3":
			fmt.Println("\nPlease wait...\n")
			CheckProductList()
		case "4":
			var inCode string
			var inPercent float32

			fmt.Printf("\nEnter Code: ")
			fmt.Scan(&inCode)
			fmt.Printf("Enter discount amount(%): ")
			fmt.Scan(&inPercent)
			CreateDiscount(inCode, inPercent)
			fmt.Println("Discount created.\n")

		case "5":
			var inCode string

			fmt.Printf("\nEnter Discount Code: ")
			fmt.Scan(&inCode)
			fmt.Println("Removing discount...")
			DeleteDiscount(inCode)
			fmt.Println("\nDiscount removed.\n")
		case "6":
			fmt.Println("\nPlease wait...\n")
			CheckDiscountList()
		case "7":
			fmt.Println("\nPlease wait...\n")
			CheckOrderList()
		case "8":
			fmt.Println("\nSigning out...\n")
			mAuth = 1
		default:
			fmt.Println("Incorrect input. Please try again.")
		}
	}
}
