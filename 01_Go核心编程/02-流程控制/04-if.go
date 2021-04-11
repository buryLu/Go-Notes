package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))

}

func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func powers(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("lim:%f < v:%f\n", lim, v)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println("start powers function")
	fmt.Println(
		powers(3, 2, 10),
		powers(3, 3, 10),
	)

	fmt.Println("start power function")
	fmt.Println(
		power(3, 2, 10),
		power(3, 3, 10),
	)

}
