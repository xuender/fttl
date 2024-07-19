package fttl_test

import (
	"fmt"

	"github.com/xuender/fttl"
)

func ExampleHash() {
	fmt.Println(fttl.Hash([]byte("test")))

	// Output:
	// 2771506328522617498
}

func ExampleIntHash() {
	fmt.Println(fttl.IntHash(1))
	fmt.Println(fttl.IntHash(2))
	fmt.Println(fttl.IntHash(3))

	// Output:
	// 3139486886484431350
	// 3238141947922148205
	// 4470094458844362118
}
