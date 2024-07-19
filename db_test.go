// nolint: errcheck
package fttl_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xuender/fttl"
)

func TestDB_Get(t *testing.T) {
	t.Parallel()

	ass := require.New(t)
	fdb := fttl.New(os.TempDir())
	key := []byte("get")

	defer fdb.Delete(key)

	fdb.PutTTL(key, []byte("value"), time.Millisecond*100, time.Millisecond*100)

	_, err := fdb.Get(key)
	ass.NoError(err)

	time.Sleep(time.Millisecond * 60)

	_, err = fdb.Get(key)
	ass.NoError(err)

	time.Sleep(time.Millisecond * 60)

	_, err = fdb.Get(key)
	ass.NoError(err)
}
