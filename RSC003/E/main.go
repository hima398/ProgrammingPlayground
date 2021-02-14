package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	I int
	J int
	D int
}

type Queue struct {
	ps []Pos
}

func NewQueue() *Queue {
	return new(Queue)
}

func (this *Queue) Push(p Pos) {
	this.ps = append(this.ps, p)
}

func (this *Queue) Pop() *Pos {
	if len(this.ps) == 0 {
		return nil
	}
	p := this.ps[0]
	this.ps = this.ps[1:]
	return &p
}

func (this *Queue) Find(i, j int) bool {
	for _, v := range this.ps {
		if i == v.I && j == v.J {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
}

var di = []int{-1, 0, 1, 0}
var dj = []int{0, -1, 0, 1}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	a := make([]string, h)
	d := make([][]int, h+1)
	for i := 0; i < h; i++ {
		a[i] = nextString()
	}
	q := NewQueue()
	for i := 0; i < h; i++ {
		d[i+1] = make([]int, w+1)
		for j := 0; j < w; j++ {
			if a[i][j] == '#' {
				d[i+1][j+1] = 0
				ok := false
				for k := 0; k < 4; k++ {
					ni, nj := i+di[k], j+dj[k]
					if ni >= 0 && nj >= 0 && ni < h && nj < w {
						ok = ok || a[ni][nj] == '.'
					}
				}
				if ok {
					p := Pos{i + 1, j + 1, 0}
					q.Push(p)
				}
			} else {
				d[i+1][j+1] = -1
			}
		}
	}
	//fmt.Println(q.Size())
	ans := 0
	for q.Size() > 0 {
		p := q.Pop()
		d[p.I][p.J] = p.D
		ans = Max(ans, p.D)
		for k := 0; k < 4; k++ {
			ni, nj := p.I+di[k], p.J+dj[k]
			nd := p.D + 1
			if ni > 0 && nj > 0 && ni <= h && nj <= w && d[ni][nj] < 0 {
				np := Pos{ni, nj, nd}
				q.Push(np)
				d[ni][nj] = nd
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
