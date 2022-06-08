package main

import "fmt"

func MapArrayData(req []int) []int {
	idx := 0
	requ := 0
	for i := 0; i < len(req); i++ {
		if requ == 0 && i != idx {
			requ = req[i]
			fmt.Println("1", requ)
		}else if i != idx {
			requ = req[i]
			fmt.Println("2", requ)
		}
		req[idx] = req[i]*requ
		fmt.Println("3", requ)

	}

	return req
}
func main() {
	req := []int{3, 27, 4, 2}
	res := MapArrayData(req)
	fmt.Println("eeeeeee", res)
}
