package fttl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/fttl"
)

func TestIntHash(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	max := 1_000_000
	set := make(map[uint64]bool, max)

	for i := range max {
		set[fttl.IntHash(i)] = true
	}

	ass.Len(set, max)
}
