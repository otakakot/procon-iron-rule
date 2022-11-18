package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	res := make([]string, 10)
	for i := 0; i < 10; i++ {
		res[i] = "0"
	}

	i := 0

	for {
		r := n % 2
		n = n / 2
		res[i] = strconv.Itoa(r)
		i = i + 1
		if n == 0 {
			break
		}
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	fmt.Println(strings.Join(res, ""))
}
