package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	count := 0
	for {
		if diff := (z*z - x) / (2 * z); math.Abs(diff) < 1.0e-15 {
			return z, count
		} else {
			z -= diff
			count++
		}
	}
}

func main() {
	z, count := Sqrt(2)
	fmt.Println("square root: ", z)
	fmt.Println("loop count: ", count)
	fmt.Println("diff with math.sqrt: ", math.Abs(z-math.Sqrt(2)))
}
