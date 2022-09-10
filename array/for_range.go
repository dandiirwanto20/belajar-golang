package main

import "fmt"

func main() {
	// cara sederhana melooping data array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// nilai tiap elemen ditampung pada variabel fruit
	// indeks titampung pada variabel i
	for i, fruit := range fruits {
		fmt.Printf("element %d : %s\n", i, fruit)
	}
}
