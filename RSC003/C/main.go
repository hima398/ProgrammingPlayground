package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	s := nextString()
	t := nextString()
	l := Lcm(n, m)
	ok := true
	gcd := Gcd(n, m)
	nn := n / gcd
	mm := m / gcd
	for i := 0; i < gcd; i++ {
		if s[i*nn] != t[i*mm] {
			ok = false
			break
		}
	}
	if ok {
		fmt.Println(l)
	} else {
		fmt.Println(-1)
	}
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

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}
