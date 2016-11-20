// CS408_Go_Project project main.go

package main

var userList []user
var productList []product
var discountList []discount

func AddUserToDatabase(u user) {
	userList = append(userList, u)
}

func DeleteUserFromDatabase(i int) {
	userList = append(userList[:i], userList[i+1:]...)
}

func GetUserLength() int {
	return len(userList)
}

func AddProductToDatabase(p product) {
	productList = append(productList, p)
}

//ProductDelete - Delete a Product from the Product list using index
func DeleteProductFromDatabase(i int) {
	productList = append(productList[:i], productList[i+1:]...)
}

func GetProductLength() int {
	return len(productList)
}

func AddDiscountToDatabase(d discount) {
	discountList = append(discountList, d)
}

func DeleteDiscountFromDatabase(i int) {
	discountList = append(discountList[:i], discountList[i+1:]...)
}

func GetDiscountLength() int {
	return len(discountList)
}
