package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuender/fttl"
)

func main() {
	fdb := fttl.New(filepath.Join(os.TempDir(), "base"))
	defer fdb.Close()

	fdb.Put([]byte("key"), []byte("value"))

	val, _ := fdb.Get([]byte("key"))
	fmt.Println(string(val))
	fmt.Println(fdb.Has([]byte("key")))
	fdb.Put([]byte("key"), []byte("new value"))

	val, _ = fdb.Get([]byte("key"))
	fmt.Println(string(val))

	fdb.Delete([]byte("key"))
}
