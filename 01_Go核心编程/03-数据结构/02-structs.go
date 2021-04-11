package main

import (
	"fmt"
)

type Age struct {
	age1 int
	age2 int
}

func main() {
	fmt.Println(Age{0, 18})

	fmt.Println("start struct fileds")
	a := Age{1, 10}
	a.age1 = 18
	fmt.Printf("a: %d, a.age1: %d\n", a, a.age1)

}
