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
