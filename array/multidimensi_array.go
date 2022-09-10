package main

import "fmt"

func main() {
	var numbers1 = [2][4]int{[4]int{3, 2, 3, 1}, [4]int{3, 4, 5, 9}} // cara pertama
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}                   // cara kedua, lebih efektif

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
}
