package main

import "fmt"

func ubahNilai(a *int) {
	*a = *a + 10 // dereference pointer untuk mengubah nilai asli
}

func main() {
	x := 5
	fmt.Println("Sebelum ubahNilai:", x)

	ubahNilai(&x) // passing alamat x ke fungsi
	fmt.Println("Setelah ubahNilai:", x)
}

// Penjelasan:
// 1. *int → artinya parameter adalah pointer ke int.
// 2. &x → mengambil alamat memori dari x.
// 3. *a = *a + 10 → akses nilai di alamat memori dan ubah.
