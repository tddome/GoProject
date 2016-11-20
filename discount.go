package main

//"fmt"

//Discount - Stores DiscountCode, DiscountPercent
type discount struct {
	dCode    string
	dPercent int
}

func GetIndexOfDiscount(code string) int {
	for i, a := range discountList {
		if a.dCode == code {
			return i
		}
	}
	return -1
}

func CreateDiscount(code string, p int) {
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

//Discount calculations below
func DiscountStuff() {

}
