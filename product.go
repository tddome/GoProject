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
func GetIndexOfProduct(id int) int {
	for i, a := range productList {
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

func DeleteProduct(id int) {
	var i int
	i = GetIndexOfProduct(id)
	DeleteProductFromDatabase(i)
}

//ProductUpdateName - Update Name from a Product using index
func ProductUpdateName(id int, n string) {
	var i int
	i = GetIndexOfProduct(id)
	productList[i].pName = n
}

//ProductUpdateID - Update ID from a Product using index
func ProductUpdateID(id int, n int) {
	var i int
	i = GetIndexOfProduct(id)
	productList[i].pID = n
}

//ProductUpdateType - Update Type from a Product using index
func ProductUpdateType(id int, n string) {
	var i int
	i = GetIndexOfProduct(id)
	productList[i].pType = n
}

//ProductUpdatePrice - Update Price from a Product using index
func ProductUpdatePrice(id int, n float32) {
	var i int
	i = GetIndexOfProduct(id)
	productList[i].pPrice = n
}

//ProductToString - Prints information on Product
func ProductToString(id int) {
	var i int
	i = GetIndexOfProduct(id)

	fmt.Println("Name:  ", productList[i].pName)
	fmt.Println("ID:       ", productList[i].pID)
	fmt.Println("Type:    ", productList[i].pType)
	fmt.Print("Price:   ")
	fmt.Printf("%6.2f", productList[i].pPrice)
	fmt.Println(" USD\n")
}
