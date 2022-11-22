package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n string
	fmt.Scan(&n)

	nn := make([]int, 0, 8)

	for _, num := range n {
		nu, _ := strconv.Atoi(string(num))
		nn = append(nn, nu)
	}

	for i := 0; i < len(nn)/2; i++ {
		nn[i], nn[len(nn)-i-1] = nn[len(nn)-i-1], nn[i]
	}

	res := 0

	for i, num := range nn {
		res = res + int(math.Pow(2, float64(i)))*num
	}

	fmt.Println(res)
}
