package main

import "fmt"

func main() {
	// khai báo biến thường
	var x int = 5748

	// khai báo và khởi tạo con trỏ với địa chỉ biến x
	var p *int = &x

	// In kết quả
	fmt.Println("Giá trị lưu trong x = ", x)
	fmt.Println("Địa chỉ của x = ", &x)
	fmt.Println("Giá trị địa chỉ lưu trong biến p = ", p)
	fmt.Println("Giá trị *p = ", *p)

	// khai báo con trỏ nhưng không khởi tạo
	var s *int

	// hiển thị giá trị của con trỏ
	fmt.Println("s = ", s)

	var y = 458
	var pointer = &y

	fmt.Println("Giá trị y = ", y)
	fmt.Println("Địa chỉ y = ", &y)
	fmt.Println("Giá trị con trỏ p = ", pointer)

	// Giải tham chiếu: truy xuất giá trị tại địa chỉ pointer
	fmt.Println("Giá trị tại địa chỉ p (*p) = ", *pointer)

	var firstName string
	// var lastName string
	// var age int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println(firstName)
}
