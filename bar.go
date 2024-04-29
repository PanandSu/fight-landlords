package main

import (
	"fight-landlords/progress"
	"fmt"
	"sort"
)

func main() {
	//check("J798K975LKK36XQAQ", "6Q538XJ6475423X6A", "48XJ2954JQ4A22A8K")
	//println(util.SortedStr("73J9579685368A4527634AA5922A4424886QJKKQQXJXQJXXKL"))
	//println(util.SortedStr("A23456789XJQKLB"))
	g()
}

func check(c1, c2, c3 string) bool {
	s := c1 + c2 + c3
	sorted := progress.AscendStr(s)
	fmt.Println(sorted)
	return true
}

func f(s string) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return s[i] < s[j]
	})
	return string(bs)
}

type Deno struct {
	Pair [2]string
}

func g() {
	var deno *Deno
	deno.Pair = [2]string{}
}
