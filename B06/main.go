package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)

	aa := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		aa[i] = a
	}

	fmt.Scan(&q)

	hits := make([]int, n+1)
	hits[0] = 0
	misses := make([]int, n+1)
	misses[0] = 0

	for i := 1; i < n; i++ {
		if aa[i-1] == 1 {
			hits[i] = hits[i-1] + 1
			misses[i] = misses[i-1]
		} else {
			hits[i] = hits[i-1]
			misses[i] = misses[i-1] + 1
		}
	}

	if aa[n-1] == 1 {
		hits[n] = hits[n-1] + 1
		misses[n] = misses[n-1]
	} else {
		hits[n] = hits[n-1]
		misses[n] = misses[n-1] + 1
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		hit := hits[r] - hits[l-1]
		miss := misses[r] - misses[l-1]
		res := hit - miss
		if res > 0 {
			fmt.Println("win")
		} else if res == 0 {
			fmt.Println("draw")
		} else {
			fmt.Println("lose")
		}
	}
}
