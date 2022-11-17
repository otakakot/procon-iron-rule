package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	aa := make([]int, n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		aa[i] = a
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || j == k || k == i {
					break
				}
				if aa[i]+aa[j]+aa[k] == 1000 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}

	fmt.Println("No")
}
