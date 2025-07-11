package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x
	for z*z >= x {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
