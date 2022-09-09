package main

import "fmt"

func main() {
	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		// fallthrough -> untuk pengecekan case dapat dilanjutkan selanjutnya
	case point > 5:
		fmt.Println("you need learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more!")
		}
	}
}
