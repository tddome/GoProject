package main

import (
	"fmt"
)

//Product - Stores Price, Name, Description
type Product struct {
	pName  string
	pID    int
	pType  string
	pPrice float32
}

//GetIndexOfProduct - Check if Product is contained by name
//Return the index where name is found, else return -1
func GetIndexOfProduct(p []*Product, id int) int {
	for i, a := range p {
		if a.pID == id {
			return i
		}
	}
	return -1
}

//ProductDelete - Delete a Product from the Product list using index
func ProductDelete(p []*Product, i int) []*Product {
	return append(p[:i], p[i+1:]...)
}

//ProductUpdatePrice - Update price from a Product using index
func ProductUpdatePrice(p []*Product, i int, v float32) []*Product {
	p[i].pPrice = v
	return p
}

//ProductUpdateName - Update name from a Product using index
func ProductUpdateName(p []*Product, i int, n string) []*Product {
	p[i].pName = n
	return p
}

//ProductUpdateDesc - Update desc from a Product using index
func ProductUpdateID(p []*Product, i int, id int) []*Product {
	p[i].pID = id
	return p
}

//ProductUpdateType - Update desc from a Product using index
func ProductUpdateType(p []*Product, i int, t string) []*Product {
	p[i].pType = t
	return p
}

//ProductToString - Prints information on Product
func ProductToString(p *Product) {
	fmt.Println("Name:  ", p.pName)
	fmt.Println("ID:       ", p.pID)
	fmt.Println("Type:    ", p.pType)
	fmt.Print("Price:   ")
	fmt.Printf("%6.2f", p.pPrice)
	fmt.Println(" USD\n")
}
