package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var x, y int
	fmt.Print("Enter x: ", x)
	fmt.Scan(&x)
	fmt.Print("Enter y: ", y)
	fmt.Scan(&y)

	fmt.Println("Before: x =", x, ", y =", y)

	swap(&x, &y)
	fmt.Println("After: x =", x, ", y =", y)
}
