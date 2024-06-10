# First Rune Function

This program defines a function that returns the first rune of a given string. A rune in Go represents a Unicode code point, which allows for handling characters from various languages and symbol sets.

## Instructions

* The function FirstRune takes a single argument, a string, and returns the first rune of that string.

## Expected Function Signature

```bash
func FirstRune(s string) rune {

}
```
## Usage

* To test the FirstRune function, you can use the following Go program:

```bash
package main

import (
	"github.com/01-edu/z01"
	"piscine"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}
```

## Expected Output

* When you run the test program, the expected output is:

```bash
$ go run .
HSO
$
```