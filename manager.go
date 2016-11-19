package main

import (
	"fmt"
)

type manager struct {
	mID string
	mPW string
}

func createProduct(p []*Product, name string, id int, atype string, price float32) []*Product {
	pNew := &Product{
		pName:  name,
		pID:    id,
		pType:  atype,
		pPrice: price,
	}

	p = append(p, pNew)
	return p
}

func checkProductList(p []*Product) {
	fmt.Println("Manager requested for list of products:\n")
	for _, prod := range p {
		ProductToString(prod)
	}
}

func checkOrderList() {
	//will work once Order is implemented
}

func managerAuthenticated(p []*Product) {
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

			fmt.Println("Creating product...")
			p = createProduct(p, "BEAR", 102, "Misc.", 1400)
			fmt.Println("Product created.")
		case "2":
			fmt.Println("Please wait...\n")
			checkProductList(p)
		case "3":
			fmt.Println("Order not implemented yet!")
		case "4":
			mAuth = 1
		default:
			fmt.Println("Incorrect input. Please try again.")
		}
	}
}
