package main

import (
	"fmt"
)

//Product - Stores Price, Name, Description
type product struct {
	pName  string
	pID    int
	pType  string
	pPrice float32
}

//GetIndexOfProduct - Check if Product is contained by name
//Return the index where name is found, else return -1
func GetIndexOfProduct(p []product, id int) int {
	for i, a := range p {
		if a.pID == id {
			return i
		}
	}
	return -1
}

//ProductCreate - Create a Product to add to the Product List
func CreateProduct(name string, id int, atype string, price float32) {
	pNew := product{
		pName:  name,
		pID:    id,
		pType:  atype,
		pPrice: price,
	}
	AddProductToDatabase(pNew)
}

//ProductDelete - Delete a Product from the Product list using index
func ProductDelete(p []product, i int) []product {
	return append(p[:i], p[i+1:]...)
}

//ProductUpdateName - Update Name from a Product using index
func ProductUpdateName(p []product, i int, n string) []product {
	p[i].pName = n
	return p
}

//ProductUpdateID - Update ID from a Product using index
func ProductUpdateID(p []product, i int, id int) []product {
	p[i].pID = id
	return p
}

//ProductUpdateType - Update Type from a Product using index
func ProductUpdateType(p []product, i int, t string) []product {
	p[i].pType = t
	return p
}

//ProductUpdatePrice - Update Price from a Product using index
func ProductUpdatePrice(p []product, i int, v float32) []product {
	p[i].pPrice = v
	return p
}

//ProductToString - Prints information on Product
func ProductToString(p product) {
	fmt.Println("Name:  ", p.pName)
	fmt.Println("ID:       ", p.pID)
	fmt.Println("Type:    ", p.pType)
	fmt.Print("Price:   ")
	fmt.Printf("%6.2f", p.pPrice)
	fmt.Println(" USD\n")
}
