package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) (res []string) {
	split := strings.Split(str, " ")
	fmt.Println("len", len(split))
	for i := len(split) - 1; i >= 0; i-- {
		res = append(res, split[i])
	}
	return res
}
func main() {
	reverse := Reverse("I am A Great human")
	fmt.Println("main", reverse)
}
