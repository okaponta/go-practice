package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	z := 1.0
	count := 0
	for {
		if diff := (z*z - x) / (2 * z); math.Abs(diff) < 1.0e-15 {
			return z, nil
		} else {
			z -= diff
			count++
		}
	}
}

func printWithErr(num float64, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}

func main() {
	printWithErr(Sqrt(2))
	printWithErr(Sqrt(-2))
}
