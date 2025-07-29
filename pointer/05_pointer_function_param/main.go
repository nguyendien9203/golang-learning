package main

import "fmt"

// Vậy tại sao không cần return?
// Vì:
// - Khi bạn truyền địa chỉ của biến num cho hàm: doubleValue(&num),
// - Trong hàm doubleValue, bạn truy cập trực tiếp địa chỉ vùng nhớ chứa num thông qua *ptr,
// - Do đó, bất kỳ thay đổi nào với *ptr sẽ ảnh hưởng trực tiếp tới giá trị của num ở main.
// => Không cần return vì bạn không tạo ra bản sao của num, bạn đang thao tác trực tiếp lên chính nó!
func doubleValue(ptr *int) {
	*ptr = *ptr * 2
}

func main() {
	var num int
	fmt.Print("Enter the number to be duplicated: ")
	fmt.Scan(&num)

	fmt.Println("Before call func: ", num)

	doubleValue(&num)

	fmt.Println("After call func: ", num)
}
