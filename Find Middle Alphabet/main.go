package main

import (
	"fmt"
	"strings"
)

func MiddleAlphabet(alpha string, l1, l2 string) (odd, even string) {
	split := strings.Split(alpha, "")
	idxL1 := 0
	idxL2 := 0
	mid := 0
	mid2 := 0
	for i, v := range split {
		if v == l1 {
			idxL1 = i
		}
		if v == l2 {
			idxL2 = i
		}
	}
	if idxL1 == 0 || idxL2 != 0 {
		AwAkh := idxL1 + idxL2
		if AwAkh%2 == 0 {
			mid = AwAkh / 2
		} else {
			mid = AwAkh / 2
			mid2 = mid + 1
		}
	}
	if mid2 == 0 {
		odd = split[mid]
	} else {
		odd = split[mid]
		even = split[mid2]
	}
	return odd, even
}

func main() {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	later1 := "Q"
	later2 := "Z"
	odd, even := MiddleAlphabet(alpha, later1, later2)

	if even == "" {
		fmt.Println(odd)
	} else {
		fmt.Println(odd+even)
	}
}
