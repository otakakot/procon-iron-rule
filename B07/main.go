package main

import "fmt"

func main() {
	var d, n int
	fmt.Scan(&d)
	fmt.Scan(&n)

	ls := make([]int, n)
	rs := make([]int, n)
	var l, r int
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r)
		ls[i] = l
		rs[i] = r
	}

	b := make([]int, d)
	for i := 0; i < n; i++ {
		b[ls[i]] += 1
		b[rs[i]] -= 1
	}

	res := make([]int, d)
	for i := 0; i < d; i++ {
		if i == 0 {
			res[i] = b[i]
		} else {
			res[i] = res[i-1] + b[i]
		}
		fmt.Println(res[i])
	}
}
