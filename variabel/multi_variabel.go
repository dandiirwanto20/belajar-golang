package main

import "fmt"

func main() {
	var first, second, third string

	first, second, third = "satu", "dua", "tiga"

	var fourt, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Printf("number: %s %s %s %s %s %s %s %s %s \n", first, second, third, fourt, fifth, sixth, seventh, eight, ninth)
}