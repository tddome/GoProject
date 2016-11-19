// GoProject

//Contributors:
// Shaylyn Wetts
// Troy Dome

// References:
// https://www.reddit.com/r/golang/comments/2uko6l/algorithms_and_data_structures_implemented_in_go/
// https://www.goinggo.net/2013/09/iterating-over-slices-in-go.html
// http://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go

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
	p1 := &Product{
		price:       8,
		name:        "Gum",
		description: "A really expensive piece of gum",
	}

	p2 := &Product{
		price:       8,
		name:        "Philly Cheese Steak",
		description: "A good sandwich",
	}

	//make a slice of products
	productList := []*Product{p1, p2}

	fmt.Println("\nList of products:")

	//array/slice loop
	//http://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go
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
		var index = GetIndexOfProduct(productList, "Gum")
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
			var index = GetIndexOfProduct(productList, "Gum")
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

	fmt.Println("\nPart 2: Shaylyn")

	fmt.Println("Hello World!")
	JamesBank := payAccount{"test", 1, 1}
	James := user{1, "password", "email", JamesBank}
	something()

	fmt.Println(James.id)
}
