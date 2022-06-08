package main

import (
	"fmt"
	"strings"
	"unicode"
)

func AlternateCase(req string) (res string) {
	resSplit := strings.Split(req, "")
	for i := 0; i < len(resSplit); i++ {
		for _, r := range resSplit[i] {
			if unicode.IsUpper(r) {
				s := strings.ToLower(resSplit[i])
				res += s
			}
			if unicode.IsLower(r) {
				s := strings.ToUpper(resSplit[i])
				res += s
			}
		}
	}
	fmt.Println("dd", res)
	return res
}
func main() {
	// res := AlternateCase("hEllO")
	// fmt.Println(res)

}
