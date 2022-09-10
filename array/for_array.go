package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		// %d = value dari array
		// %s = isi/nama element dari fruits
		fmt.Printf("element %d : %s\n", i, fruits[i])
	}

	// contoh perulangan hanya kondisi
	// var i = 0

	// for i < len(fruits) {
	// 	fmt.Printf("element %d : %s\n", i, fruits[i])
	// 	i++
	// }

	// contoh perulangan tanpa argumen
	// var i = 0

	// for {
	// 	fmt.Printf("element %d : %s\n", i, fruits[i])
	// 	i++
	// 	if i == len(fruits) {
	// 		break
	// 	}
	// }
}
