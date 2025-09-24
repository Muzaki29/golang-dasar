package main

import "fmt"

// variadic function: bisa menerima banyak angka
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// memanggil fungsi variadic dengan beberapa argumen
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", result)

	// menggunakan defer untuk menunda eksekusi fungsi
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
}
