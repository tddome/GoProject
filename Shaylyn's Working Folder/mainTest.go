// CS408_Go_Project project main.go
package mainTest

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	JamesBank := payAccount{"test", 1, 1}
	James := user{1, "password", "email", JamesBank}
	something()

	fmt.Println(James.id)
}
