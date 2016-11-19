package main

import (
	"fmt"
)

type manager struct {
	mID string
	mPW string
}

func createProduct(p []*Product, name string, id int, atype string, price float32) []*Product {
	pNew := &Product{
		pName:  name,
		pID:    id,
		pType:  atype,
		pPrice: price,
	}

	p = append(p, pNew)
	return p
}

func checkProductList(p []*Product) {
	fmt.Println("Manager requested for list of products:\n")
	for _, prod := range p {
		ProductToString(prod)
	}
}

func checkOrderList() {
	//will work once Order is implemented
}
