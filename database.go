// CS408_Go_Project project main.go

package main

var userList []user
var productList []product

func AddUserToDatabase(u user) {
	userList = append(userList, u)
}

func ReturnUserListLength() int {
	return len(userList)
}

func AddProductToDatabase(p product) {
	productList = append(productList, p)
}

func ReturnProductListLength() int {
	return len(productList)
}
