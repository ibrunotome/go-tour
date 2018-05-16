package main

import "fmt"

func addWithSameType(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(addWithSameType(42, 13))
}
