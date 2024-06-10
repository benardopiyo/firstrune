package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	run := []rune(s)
	return run[0]
}

func main() {
	z01.PrintRune(FirstRune("Hello"))
	z01.PrintRune(10)
}
