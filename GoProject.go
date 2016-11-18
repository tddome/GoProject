// GoProject

//Contributors:
// Shaylyn Wetts
// Troy Dome

// References:
// https://www.reddit.com/r/golang/comments/2uko6l/algorithms_and_data_structures_implemented_in_go/
// https://www.goinggo.net/2013/09/iterating-over-slices-in-go.html

package main

import (
	"bufio"
	"fmt"
	"os"
)

//Product - Stores Price, Name, Description
type Product struct {
	price       float32
	name        string
	description string
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
	//http://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Is the gum too expensive for you? y/n: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	if text == "y" {
		//https://www.reddit.com/r/golang/comments/283vpk/help_with_slices_and_passbyreference/
		fmt.Println("\nJust run please")
	}

	fmt.Println("List of products:")

	for _, prod := range productList {
		ProductToString(prod)
	}

	fmt.Println("Part 2: Shaylyn")

	fmt.Println("Hello World!")
	JamesBank := payAccount{"test", 1, 1}
	James := user{1, "password", "email", JamesBank}
	something()

	fmt.Println(James.id)
}
