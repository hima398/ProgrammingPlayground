package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// https://twitter.com/RinSakamichi/status/1359812034100162561
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ans := 0
	for i := 0; i < n; i++ {
		a := nextString()
		if a == "R" {
			ans++
		} else if a == "B" {
			continue
		} else {
			ia, _ := strconv.Atoi(a)
			if ia%2 == 1 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
