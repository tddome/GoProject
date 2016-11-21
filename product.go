/* GoProject
*
* Class:
*	CS 408
*
* Contributors:
* 	Troy Dome
* 	Shaylyn Wetts
*
* Last Updated:
*	11/20/2016
*
* Function:
*	Product struct along with functions relating to product.
 */

package main

import (
	"fmt"
)

// Product - Stores name, ID number, type, and price
type product struct {
	pName  string
	pID    int
	pType  string
	pPrice float32
}

// GetIndexOfProduct - Check if a product is contained by ID
// and returns the index where the ID is found, else return -1
func GetIndexOfProduct(id int) int {
	for i, a := range productList {
		if a.pID == id {
			return i
		}
	}
	return -1
}

// ProductCreate - Create a product and add to database.  Checks to
// make sure the ID of the product is unique
func CreateProduct(name string, id int, atype string, price float32) {
	var count int = 0

	for count < len(productList) {
		if productList[count].pID == id {
			fmt.Println("That product id is taken.  Please choose a new ID.")
			return
		} else {
			count++
		}
	}

	pNew := product{
		pName:  name,
		pID:    id,
		pType:  atype,
		pPrice: price,
	}
	AddProductToDatabase(pNew)
}

// DeleteProduct - Delete a product from the database
func DeleteProduct(id int) {
	var i int
	i = GetIndexOfProduct(id)
	DeleteProductFromDatabase(i)
}

// ProductUpdateName - Update the name of a product
func ProductUpdateName(id int, n string) {
	var i int
	i = GetIndexOfProduct(id)
	productList[i].pName = n
}

// ProductUpdateType - Update the type of a product
func ProductUpdateType(id int, n string) {
	var i int
	i = GetIndexOfProduct(id)
	productList[i].pType = n
}

// ProductUpdatePrice - Update the price of a product
func ProductUpdatePrice(id int, n float32) {
	var i int
	i = GetIndexOfProduct(id)
	productList[i].pPrice = n
}

// GetProductPrice - Returns the price of a product
func GetProductPrice(id int) float32 {
	var i int
	i = GetIndexOfProduct(id)
	return productList[i].pPrice
}

// ProductToString - Prints information of a product
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
