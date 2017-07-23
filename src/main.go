package main

import "fmt"

func divMod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	var a, b int

	a, b = divMod(7, 12)
	fmt.Println("divMod(7, 12) =", a, b)

	div, mod := divMod(36, 12)
	fmt.Println("divMod(36, 12) =", div, mod)

	_, i127mod12 := divMod(127, 12)
	fmt.Println("127 modulo 12 =", i127mod12)

}
