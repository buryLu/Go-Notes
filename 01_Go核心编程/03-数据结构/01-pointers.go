package main

import (
	"fmt"
)

func main() {
	i, j := 10, 20

	p := &i
	fmt.Println(p)

	*p = 100
	fmt.Println(i)

	m := &p
	fmt.Println(m)

	**m = 520

	fmt.Printf("m: %p,P: %p,I: %d\n", m, p, i)

	p = &j
	fmt.Println(p)

	*p = 200
	fmt.Println(j)
}
