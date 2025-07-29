package main

import "fmt"

func newInt(value int) *int {
	p := new(int) // Cấp phát động một biến int, p là con trỏ đến vùng nhớ mới
	*p = value    // Gán giá trị được truyền vào cho vùng nhớ đó
	return p      // Trả về con trỏ đó
}

func main() {
	var v int
	fmt.Print("Enter the value to allocate: ")
	fmt.Scan(&v)

	p := newInt(v)
	fmt.Println("The value at pointer p: ", *p)
}
