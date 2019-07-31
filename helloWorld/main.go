package main

import "fmt"

func main() {

	firstName := "Aurthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)

}


// What I did here was - stored a variable in memory called firstName -- which held string data "Aurthur"

// Then in that SAME memory address I changed firstName to Tricia

