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
*	Creates a manager.  Includes functions used by the
*	manager for adding products, removing products, adding
*	discounts, removing discounts, and accessing all current
*	products, discounts, and orders.  Includes prompts for
*	user input to run these functions.
 */

package main

import (
	"fmt"
	"time"
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
			fmt.Println("\nEnter Name: Bear")
			fmt.Println("Enter ID: 99")
			fmt.Println("Enter Type: Animal")
			fmt.Println("Enter Price: 1000000")
			fmt.Println("\nCreating product...")
			time.Sleep(2 * time.Second)
			CreateProduct("BEAR", 99, "ANIMAL", 1000000)
			fmt.Println("Product created.\n")

			// Extra products for the sake of the demo
			CreateProduct("PAINT BRUSH", 406, "ART SUPPLIES", 15)
			CreateProduct("PHILLY CHEESE STEAK", 101, "FOOD", 7)
			CreateProduct("SANDALS", 51, "CLOTHING", 20)
			CreateProduct("DOG TOY DELUXE", 203, "TOY", 30)
			CreateProduct("CLASSIC TANK TOP", 309, "CLOTHING", 90)
		case "2":
			fmt.Println("\nEnter Product ID: 51")
			fmt.Println("Removing product...\n")
			time.Sleep(2 * time.Second)
			DeleteProduct(51)
			fmt.Println("Product removed.\n")
		case "3":
			fmt.Println("\nPlease wait...\n")
			time.Sleep(2 * time.Second)
			CheckProductList()
		case "4":
			fmt.Println("\nEnter Code: SAVE20")
			fmt.Println("Enter discount amount(%): 20")
			time.Sleep(2 * time.Second)
			CreateDiscount("SAVE20", 20)
			fmt.Println("Discount created.\n")

			// Extra discounts for the sake of the demo
			CreateDiscount("WOWMAN", 70)
			CreateDiscount("BEARSPECIAL", 40)
			CreateDiscount("SWEET30", 30)
			CreateDiscount("CHRISTMASJOY", 50)
		case "5":
			fmt.Println("\nEnter Discount Code: CHRISTMASJOY")
			fmt.Println("Removing discount...")
			time.Sleep(2 * time.Second)
			DeleteDiscount("CHRISTMASJOY")
			fmt.Println("\nDiscount removed.\n")
		case "6":
			fmt.Println("\nPlease wait...\n")
			time.Sleep(2 * time.Second)
			CheckDiscountList()
		case "7":
			fmt.Println("\nPlease wait...\n")
			time.Sleep(2 * time.Second)
			CheckOrderList()
		case "8":
			fmt.Println("\nSigning out...\n")
			mAuth = 1
		default:
			fmt.Println("Incorrect input. Please try again.")
		}
	}
}
