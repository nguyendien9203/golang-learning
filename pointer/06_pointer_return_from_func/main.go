package main

import "fmt"

func createPointer() *int {
	x := 50
	return &x
}

func main() {
	p := createPointer()
	fmt.Println("The value at pointer p: ", *p)
}
