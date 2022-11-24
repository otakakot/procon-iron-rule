package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			z := k - i - j
			if z >= 1 && z <= n {
				res += 1
			}
		}
	}

	fmt.Println(res)
}
