package main

import "fmt"

const conferenceName = "Go Conference"

func main() {
	fmt.Println("Welcome to our", conferenceName, "booking application!")
	fmt.Println("Thank you for booking a ticket for the", conferenceName)
	fmt.Println("Printing ticket for", conferenceName, "…")

	name := "Dien"
	age := 25
	height := 1.75
	member := true

	// In thông tin với định dạng
	fmt.Printf("Name: %s\n", name)              // %s dùng để in chuỗi (string)
	fmt.Printf("Age: %d\n", age)                // %d dùng để in số nguyên (decimal)
	fmt.Printf("Height: %.2f meters\n", height) // %f dùng để in số thực (float), %.2f giữ 2 số sau dấu thập phân
	fmt.Printf("Is Member: %t\n", member)       // %t dùng để in giá trị boolean

	// Kết hợp nhiều verb trong một chuỗi
	fmt.Printf("Hi %s, you are %d years old and your height is %.1f meters.\n", name, age, height)

	fmt.Printf("Name: %s\n", name)

	var rawString = `Hello`
	fmt.Print(rawString)
}
