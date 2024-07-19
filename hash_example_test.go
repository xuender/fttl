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
