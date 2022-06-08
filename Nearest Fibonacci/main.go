package main

import "fmt"

func fibbonaci(int) int {
	angka := 0
	angka2 := 1
	for i := 0; i < 2; i++ {
		temp := angka
		angka = angka2
		angka2 = temp + angka
	}
	return angka
}
func nearest(arr []int) int {
	var jumlah int
	var angkaFibo int
	for _, v := range arr {
		jumlah += v
	}

	var i int
	for jumlah > fibbonaci(i) {
		angkaFibo = fibbonaci(i)
		i++
	}

	nearest := angkaFibo - jumlah
	fmt.Println("j", jumlah)
	fmt.Println("f", angkaFibo)
	return nearest
}
func main() {
	fibbonaci := nearest([]int{15, 1, 3})
	fmt.Println(fibbonaci)
}
