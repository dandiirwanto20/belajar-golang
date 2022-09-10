package main

import "fmt"

func main() {
	// cara horizontal
	// fruits = [4]string{"apple", "grape", "banana", "melom"}

	// cara vertical
	var fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t \t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
}
