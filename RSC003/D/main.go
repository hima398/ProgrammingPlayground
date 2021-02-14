package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const p = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := nextString()
	m := make(map[byte]int)
	for i := 0; i < n; i++ {
		m[s[i]]++
	}
	//fmt.Println(m)
	ans := 1
	for _, v := range m {
		ans *= (v + 1)
		ans %= p
	}
	ans--
	if ans < 0 {
		ans += p
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
