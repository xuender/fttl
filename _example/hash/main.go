package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/xuender/fttl"
)

func main() {
	max := 1_000_000_000
	path := "/tmp/hash.txt"
	file, _ := os.Create(path)

	for i := range max {
		fmt.Fprintf(file, "%d\n", fttl.IntHash(i))
	}

	file.Close()

	wc(path)

	sortCmd := exec.Command("sort", path, ">", path+".sort")
	sortCmd.Stdout = os.Stdout

	sortCmd.Run()

	uniqCmd := exec.Command("uniq", path+".sort", ">", path)
	uniqCmd.Stdout = os.Stdout

	uniqCmd.Run()

	wc(path)
}

func wc(path string) {
	cmd := exec.Command("wc", "-l", path)
	cmd.Stdout = os.Stdout

	cmd.Run()
}
