package main

import (
	"fmt"
	"sort"
)

func splitToDigits(n int) []int {
	if n < 0 || n == 0 {
		return nil
	}
	var ret []int
	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}
	if len(ret) != 3 {
		return nil
	}
	return ret
}

func main() {
	n := splitToDigits(2321)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
