package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const limit = float64(1e-10)
	z := float64(1.0)
	for {
		last := z
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("z = %.15f, last = %.15f\n", z, last)
		if math.Abs(z-last) < limit {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	//fmt.Println(math.Sqrt(2))
}
