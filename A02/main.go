package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	aa := make([]int, n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		aa[i] = a
	}

	for _, a := range aa {
		if a == x {
			fmt.Print("Yes")
			return
		}
	}

	fmt.Print("No")
}
