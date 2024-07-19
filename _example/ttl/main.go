package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xuender/fttl"
)

func main() {
	fdb := fttl.New(filepath.Join(os.TempDir(), "ttl"))
	defer fdb.Close()

	fdb.PutTTL([]byte("key"), []byte("value"), time.Second, time.Second)

	time.Sleep(time.Second)

	_, err := fdb.Get([]byte("key"))
	fmt.Println(err)
}
