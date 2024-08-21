package fttl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/fttl"
)

func TestIntHash(t *testing.T) {
	t.Parallel()

	var max uint32 = 1_000_000

	ass := assert.New(t)
	set := make(map[uint64]bool, max)

	for i := range max {
		set[fttl.IntHash(i)] = true
		set[fttl.IntHash(^uint32(0)-i)] = true
	}

	ass.Len(set, int(max*2))
}
