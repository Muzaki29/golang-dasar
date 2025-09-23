package main

import "fmt"

func main() {
	angka := 10

	// IF ELSE
	if angka > 0 {
		fmt.Println("Angka positif")
	} else if angka < 0 {
		fmt.Println("Angka negatif")
	} else {
		fmt.Println("Angka nol")
	}

	// SWITCH CASE
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("Hari ini adalah hari Senin")
	case "Sabtu", "Minggu":
		fmt.Println("Hari ini adalah akhir pekan")
	default:
		fmt.Println("Hari ini adalah hari kerja")
	}

}
