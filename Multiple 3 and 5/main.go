package main

import "fmt"

func Multiple3and5(len int) (res int) {
	for i := 0; i < len; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}
	return res
}

func main() {
	res := Multiple3and5(20)
	fmt.Println(res)
}