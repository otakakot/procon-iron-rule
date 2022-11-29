package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	aa := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		aa[i] = a
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		res := 0
		for j := l - 1; j < r; j++ {
			res += aa[j]
		}
		fmt.Println(res)
	}
}
