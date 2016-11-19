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
)

//Product - Stores Price, Name, Description
type Product struct {
	price       float32
	name        string
	description string
}

//GetIndexOfProduct - Check if Product is contained by name
//Return the index where name is found, else return -1
func GetIndexOfProduct(p []*Product, n string) int {
	for i, a := range p {
		if a.name == n {
			return i
		}
	}
	return -1
}

//ProductDelete - Delete a value from a Product using index
func ProductDelete(p []*Product, i int) []*Product {
	return append(p[:i], p[i+1:]...)
}

//ProductUpdatePrice - Update a value from a Product using index
func ProductUpdatePrice(p []*Product, i int, v float32) []*Product {
	p[i].price = v
	return p
}

//ProductToString - Prints information on Product
func ProductToString(p *Product) {
	fmt.Println("Info on product: ")
	fmt.Println("Name:  ", p.name)
	fmt.Print("Price:   ")
	fmt.Printf(" $%6.2f", p.price)
	fmt.Print(" Dollar\n")
	fmt.Println("Info:    ", p.description)
}

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
		case "n":
			fmt.Println("\nAlright, keep it overpriced I guess...\n")
		default:
			fmt.Println("\nAnswer not yes or no, ignoring...\n")
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
