package main

import "fmt"

func main() {
	// contoh deklarasi konstanta
	const firstName string = "dan"

	// contoh deklarasi konstanta dengan type inference
	const lastName = "di"
	fmt.Print("Halo ", firstName, "!\n")
	fmt.Print("Nice to meet u ", lastName, "!\n")
	fmt.Println("John Wick")
	fmt.Println("John", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}