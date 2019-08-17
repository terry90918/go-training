package foo

import "fmt"

func init() {
	fmt.Println("func init!")
}

func HelloBar() string {
	return "Hello bar!"
}
