package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s []byte
	for i := 0; i < 1000; i++ {
		s = append(s, '.')
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < 1000; i++ {
		fmt.Fprintln(out, string(s))
	}
}
