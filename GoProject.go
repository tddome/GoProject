// GoProject

//Contributors:
// Shaylyn Wetts
// Troy Dome

// References:
// https://www.reddit.com/r/golang/comments/2uko6l/algorithms_and_data_structures_implemented_in_go/
// https://www.goinggo.net/2013/09/iterating-over-slices-in-go.html
// http://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go

// Project Outline (Might only work for Troy?):
// https://docs.google.com/document/d/1ocq3TYRU3JASzWDYsycaDi9CGuoG2MVsmT2LDAnvnaE/edit?ts=580ed160

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	//& used to assure new address generated for each product
	//avoids conflicts of possible over-writing in memory
	//or duplicate copies in memory being generated
	p1 := product{
		pName:  "Gum",
		pID:    100,
		pType:  "Food",
		pPrice: 8,
	}

	p2 := product{
		pName:  "Philly Cheese Steak",
		pID:    101,
		pType:  "Food",
		pPrice: 8,
	}

	//make a slice of products
	productList := []product{p1, p2}

	fmt.Println("\nList of products:")

	//array/slice loop
	for _, prod := range productList {
		ProductToString(prod)
	}

	//console input/output
	fmt.Print("Do you want to delete the expensive gum? y/n: ")
	var text string
	fmt.Scan(&text)

	switch text {
	case "y":
		fmt.Println("\nDeleting Gum...")
		var index = GetIndexOfProduct(productList, 100)
		productList = ProductDelete(productList, index)
		fmt.Println("\nGum was deleted.\n")
		time.Sleep(2 * time.Second)
	case "n":
		fmt.Println("\nNot Deleting Gum...")
		fmt.Print("Would you like to update the price of Gum to $3? y/n: ")
		fmt.Scan(&text)
		switch text {
		case "y":
			fmt.Println("\nUpdating Gum price to $3...")
			var index = GetIndexOfProduct(productList, 100)
			productList = ProductUpdatePrice(productList, index, 3)
			fmt.Println("\nPrice updated.\n")
			time.Sleep(2 * time.Second)
		case "n":
			fmt.Println("\nAlright, keep it overpriced I guess...\n")
			time.Sleep(2 * time.Second)
		default:
			fmt.Println("\nAnswer not yes or no, ignoring...\n")
			time.Sleep(2 * time.Second)
		}
	default:
		fmt.Println("\nAnswer not yes or no, ignoring...\n")
	}

	fmt.Println("List of products:")

	for _, prod := range productList {
		ProductToString(prod)
	}

	fmt.Print("Are you a manager? y/n: ")
	fmt.Scan(&text)

	switch text {
	case "y":
		fmt.Print("\nWhat is your Manager ID?: ")
		fmt.Scan(&text)
		switch text {
		case "TroyBoy420":
			fmt.Print("\nWhat is your password, TroyBoy420?: ")
			fmt.Scan(&text)
			switch text {
			case "bears":
				managerAuthenticated(productList)
			default:
				fmt.Println("\nYou're not Troy. Go away.\n")
			}
		default:
			fmt.Println("\nManager ID doesn't exist. Moving on...\n")
		}
	case "n":
		fmt.Println("\nNo worries. Proceeding to Part 2...\n")
	default:
		fmt.Println("\nAnswer not yes or no, ignoring...\n")
	}

	fmt.Println("Part 2: Shaylyn")

	fmt.Println("Hello World!")
	JamesBank := payAccount{"test", 1, 1}
	James := user{1, "password", "email", JamesBank}
	something()

	fmt.Println(James.id)
}
