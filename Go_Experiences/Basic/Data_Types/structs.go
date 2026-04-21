// structs.go

package main

import (
	"fmt"
)

type Color struct {
	R, G, B int
}

// ToHex converts the Color struct values into a CSS-style hexadecimal string.
func (c Color) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func (c Color) String() string {
	return c.ToHex()
}

func main() {
	color := Color{R: 106, G: 90, B: 205}

	fmt.Printf("RGB Values: %d, %d, %d\n", color.R, color.G, color.B)
	fmt.Println("CSS Hex format:", color.ToHex())

	fmt.Println("Direct print (Stringer):", color)
}
