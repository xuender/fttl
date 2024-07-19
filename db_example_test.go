// nolint: errcheck
package fttl_test

import (
	"fmt"
	"os"
	"time"

	"github.com/xuender/fttl"
)

func ExampleDB_Get() {
	fdb := fttl.New(os.TempDir())
	key := []byte("get")

	defer fdb.Delete(key)
	fmt.Println(fdb.Put(key, []byte("value")))

	val, err := fdb.Get(key)
	fmt.Println(err)
	fmt.Println(string(val))

	// Output:
	// <nil>
	// <nil>
	// value
}

func ExampleDB_Put() {
	fdb := fttl.New(os.TempDir())
	key := []byte("put")

	defer fdb.Delete(key)
	fdb.Put(key, []byte("value"))
	fdb.Put(key, []byte("value2"))

	val, _ := fdb.Get(key)
	fmt.Println(string(val))

	// Output:
	// value2
}

func ExampleDB_PutTTL() {
	fdb := fttl.New(os.TempDir())
	key := []byte("ttl")

	defer fdb.Delete(key)

	fdb.PutTTL(key, []byte("value"), time.Millisecond*100, 0)

	val, _ := fdb.Get(key)
	fmt.Println(string(val))
	time.Sleep(time.Millisecond * 200)

	_, err := fdb.Get(key)
	fmt.Println(err)

	// Output:
	// value
	// not found
}
