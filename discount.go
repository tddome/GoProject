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
*	Creates a discount.  Provides functions for adding,
*	removing, and printing discounts.  Also includes a
*	function for calculating a total after a discount
*	is applied.
 */

package main

import (
	"fmt"
)

//Discount - Stores DiscountCode, DiscountPercent
type discount struct {
	dCode    string
	dPercent float32
}

// GetIndexOfDiscount - Check if a discount is contained by code
// and returns the index where the code is found, else return -1
func GetIndexOfDiscount(code string) int {
	for i, a := range discountList {
		if a.dCode == code {
			return i
		}
	}
	return -1
}

// CreateDiscount - Creates a discount and adds it to
// the database
func CreateDiscount(code string, p float32) {
	dNew := discount{
		dCode:    code,
		dPercent: p,
	}
	AddDiscountToDatabase(dNew)
}

// DeleteDiscount - Deletes a current discount from the
// database
func DeleteDiscount(id string) {
	var i int
	i = GetIndexOfDiscount(id)
	DeleteDiscountFromDatabase(i)
}

//	DiscountAmount - Applies amount being discounted from total.
//	Returns the amount being discounted.
//	(So you don't, say, discount the discounted price, which is wrong)
//	(Unless you want them to save even more money)
func DiscountAmount(discountTotal float32, discountCode string) float32 {
	var indexDiscount = GetIndexOfDiscount(discountCode)
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

// DiscountToString - Prints a discount's code and amount
func DiscountToString(id string) {
	var i int
	i = GetIndexOfDiscount(id)

	fmt.Println("Discount Code: ", discountList[i].dCode)
	fmt.Println("Discount Ammount(%): ", discountList[i].dPercent)
	fmt.Println("\n")
}
