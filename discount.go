package main

import (
	"fmt"
)

//Discount - Stores DiscountCode, DiscountPercent
type discount struct {
	dCode    string
	dPercent float32
}

func GetIndexOfDiscount(code string) int {
	for i, a := range discountList {
		if a.dCode == code {
			return i
		}
	}
	return -1
}

func CreateDiscount(code string, p float32) {
	dNew := discount{
		dCode:    code,
		dPercent: p,
	}
	AddDiscountToDatabase(dNew)
}

func DeleteDiscount(id string) {
	var i int
	i = GetIndexOfDiscount(id)
	DeleteDiscountFromDatabase(i)
}

//DiscountSearch - Searches the database for index of discount
//Uses the name of the discount, returns -1 if not indexed in database
func DiscountSearch(discountCode string) int {
	for i, dis := range discountList {
		if dis.dCode == discountCode {
			return i
		}
	}
	return -1
}

//DiscountAmount - Applies amount being discounted from total
//Returns the amount being discounted
//(So you don't, say, discount the discounted price, which is wrong)
//(Unless you want them to save even more money)
func DiscountAmount(discountTotal float32, discountCode string) float32 {
	var indexDiscount = DiscountSearch(discountCode)
	if indexDiscount == -1 {
		fmt.Println("Discount not in database; No discount applied.")
		return 0
	} else {
		var discountPercent = discountList[indexDiscount].dPercent / 100
		discountPercent = 1 - discountPercent
		var originalPrice = discountTotal
		discountTotal = discountTotal * discountPercent
		fmt.Println("Discount applied.")
		return originalPrice - discountTotal
	}
}
