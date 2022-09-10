package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// contoh implementasi nested for hanya kondisi
	// var i = 0

	// for i < 5 {
	// 	var j = i

	// 	for j < 5 {
	// 		fmt.Print(j, " ")
	// 		j++
	// 	}
	// 	i++
	// 	fmt.Println()
	// }

	// contoh nested for tanpa argumen
	// var i = 0

	// for {
	// 	var j = i

	// 	for {
	// 		fmt.Print(j, " ")
	// 		j++
	// 		if j == 5 {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println()
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }
}
