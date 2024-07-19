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

	fdb.Delete([]byte("key"))
}
