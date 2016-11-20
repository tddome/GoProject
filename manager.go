package main

import (
	"fmt"
	"time"
)

type manager struct {
	mID string
	mPW string
}

func checkProductList() {
	fmt.Println("Manager requested for list of products:\n")
	for _, prod := range productList {
		ProductToString(prod.pID)
	}
}

func checkDiscountList() {
	fmt.Println("Manager requested for list of discount codes:\n")
	for _, dis := range discountList {
		DiscountToString(dis.dCode)
	}
}

func checkOrderList() {
	fmt.Println("Manager requested for list of past orders:\n")
	for _, ord := range orderHistory {
		OrderToString(ord.oBillID)
	}
}

func managerAuthenticated() {
	fmt.Println("\nWelcome, manager! What do you want to do?")
	var mAuth = 0

	for mAuth < 1 {
		fmt.Println("1. Add a product")
		fmt.Println("2. Delete a product")
		fmt.Println("3. Access product list")
		fmt.Println("4. Add a discount code")
		fmt.Println("5. Remove a discount code")
		fmt.Println("6. Access discount list")
		fmt.Println("7. Access order list")
		fmt.Println("8. Sign out")
		var text string
		fmt.Scan(&text)

		switch text {
		case "1":
			fmt.Println("Enter Name: Bear")
			fmt.Println("Enter ID: 99")
			fmt.Println("Enter Type: Animal")
			fmt.Println("Enter Price: 1000000")
			fmt.Println("\nCreating product...")
			time.Sleep(2 * time.Second)
			CreateProduct("BEAR", 99, "Animal", 1000000)
			fmt.Println("Product created.\n")
		case "2":
			fmt.Println("Enter Product ID: 69")
			fmt.Println("Removing product...\n")
			DeleteProduct(69)
			fmt.Println("Product removed.\n")
		case "3":
			fmt.Println("Please wait...\n")
			checkProductList()
		case "4":
			fmt.Println("Enter Code: SAVE20")
			fmt.Println("Enter discount amount: 0.20")
			CreateDiscount("SAVE20", 0.20)
			fmt.Println("Discount created.\n")
		case "5":
			fmt.Println("Enter Discount Code: SAVE20")
			fmt.Println("Removing discount...")
			DeleteDiscount("SAVE20")
			fmt.Println("Discount removed.\n")
		case "6":
			fmt.Println("Please wait...\n")
			checkDiscountList()
		case "7":
			fmt.Println("Please wait...\n")
			checkOrderList()
		case "8":
			fmt.Println("Signing out...\n")
			mAuth = 1
		default:
			fmt.Println("Incorrect input. Please try again.")
		}
	}
}
