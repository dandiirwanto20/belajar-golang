package main

import "fmt"

func main() {
	var numbers = [...]int{2, 3, 2, 4, 3}
	// numbers := [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah element \t:", len(numbers))
}
