package main

import (
	"fmt"

	"github.com/Luke-Vear/anybase"
)

func main() {
	base56Alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

	gen, _ := anybase.NewGenerator(base56Alphabet)
	num := gen.NumberFromUint(1<<64 - 1)

	fmt.Println(num) // 36psTsTZwTUH
}
