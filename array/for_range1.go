package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// jika dibutuhkan hanya elemen atau isinya
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// jika dibutuhkan hanya indeksnya
	// for i, _ := range fruits {
	// 	fmt.Printf("Indeks elemen : %d\n", i)
	// }
}
