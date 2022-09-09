package main

import "fmt"

func main() {
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)
	
	leftOrRight := left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)
	
	leftReverse := !left
	fmt.Printf("!left \t(%t) \n", leftReverse)

}