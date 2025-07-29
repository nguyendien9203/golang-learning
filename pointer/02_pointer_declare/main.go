package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter the value of a: ")
	fmt.Scan(&a)

	p := &a

	fmt.Println("The value of a: ", a)
	fmt.Println("Address of a: ", &a)
	fmt.Println("The value of pointer p: ", p)
}
