package main

import (
	"fmt"
)

func Sqrt(x float64) (z float64, count uint) {
	z = 1.0
	count = 0
	for d := 1.0; d*d > 1e-10; z -= d {
		d = (z*z - x) / (2 * z)
		count++
	}
	return
}

func main() {
	for x := 2.0; x <= 20.0; x++ {
		z, count := Sqrt(x)
		fmt.Println(x, z, count)
	}
}
