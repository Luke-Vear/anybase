package main

import (
	"fmt"

	"github.com/Luke-Vear/anybase"
)

func main() {
	base56Alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

	encoded, err := anybase.Encode(^uint(0), base56Alphabet)
	if err != nil {
		panic(err)
	}

	fmt.Println(encoded) // 36psTsTZwTUH
}
