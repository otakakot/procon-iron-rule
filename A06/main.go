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

	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + aa[i-1]
	}
	res[n] = res[n-1] + aa[n-1]

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Println(res[r] - res[l-1])
	}
}
