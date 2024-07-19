package fttl_test

import (
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/fttl"
)

// nolint: paralleltest
func TestDB_clean(t *testing.T) {
	ass := assert.New(t)
	ticker := time.NewTicker(time.Millisecond * 100)

	patches := gomonkey.ApplyFuncReturn(time.NewTicker, ticker)
	defer patches.Reset()

	fdb := fttl.New(os.TempDir())
	key := []byte("clean")

	_ = fdb.PutTTL(key, []byte("data"), time.Millisecond*50, time.Millisecond*50)
	time.Sleep(time.Millisecond * 120)

	_, err := fdb.Get(key)
	ass.Error(err)
}
