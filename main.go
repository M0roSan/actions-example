package main

import "fmt"

func main() {
	fmt.Println("HELLO FOO BAR")
	fmt.Printf("1 + 2 = %d", sum(1, 2))
}

func sum(n1, n2 int) int {
	return n1 + n2
}
