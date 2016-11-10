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

//object created - a prototype of Product
type Product struct {
	price       float32
	name        string
	description string
}

//prints Product 'structs' that are passed in
func ProductToString(p *Product) {
	fmt.Println("Info on product: ")
	fmt.Println("Name:  ", p.name)
	fmt.Print("Price:   ")
	fmt.Printf(" $%6.2f", p.price)
	fmt.Print(" Dollar\n")
	fmt.Println("Info:    ", p.description)
}

func main() {
	fmt.Println("Hello World!\n")

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

	//a slice, which is more popular in go than lists
	//since this seems to be a dynamically allocating array
	productList := []*Product{p1, p2}

	fmt.Println("List of products:\n")

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

	if text == "y" {
		fmt.Println("\nHold on there, class just ended.")
		fmt.Println("I was working with lists and now they're slices")
		fmt.Println("I'll change this soon so it updates to 1 dollar\n")
	}

	fmt.Println("List of products:\n")

	for _, prod := range productList {
		ProductToString(prod)
	}
}
