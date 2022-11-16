package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		if 100%i == 0 {
			fmt.Print("Yes")
			return
		}
	}

	fmt.Printf("No")
}
