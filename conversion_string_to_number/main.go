package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "100"

	angka, err := strconv.Atoi(text) // string → int
	if err != nil {
		fmt.Println("Pesan Error :", err.Error())
	} else {
		fmt.Println("Angka :", angka)
	}
}
