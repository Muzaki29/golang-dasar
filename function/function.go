package main

import "fmt"

// Fungsi dengan parameter dan return
func tambah(a int, b int) int {
	return a + b // mengembalikan hasil penjumlahan
}

func main() {
	hasil := tambah(3, 5) // memanggil fungsi tambah
	fmt.Println("Hasil penjumlahan:", hasil)
}
