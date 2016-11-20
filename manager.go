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

func checkOrderList() {
	//will work once Order is implemented
}

func managerAuthenticated(p []product) {
	fmt.Println("\nWelcome, manager! What do you want to do?")
	var mAuth = 0

	for mAuth < 1 {
		fmt.Println("1. Add a product")
		fmt.Println("2. Access product list")
		fmt.Println("3. Access order list")
		fmt.Println("4. Sign out")
		var text string
		fmt.Scan(&text)

		switch text {
		case "1":
			fmt.Println("Enter Name: Bear")
			fmt.Println("Enter ID: 69")
			fmt.Println("Enter Type: Animal")
			fmt.Println("Enter Price: 420")
			fmt.Println("\nCreating product...")
			time.Sleep(2 * time.Second)
			CreateProduct("BEAR", 69, "Animal", 420)
			fmt.Println("Product created.\n")
		case "2":
			fmt.Println("Please wait...\n")
			checkProductList()
		case "3":
			fmt.Println("Order not implemented yet!")
		case "4":
			fmt.Println("Signing out...\n")
			mAuth = 1
		default:
			fmt.Println("Incorrect input. Please try again.")
		}
	}
}
