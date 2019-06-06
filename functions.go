package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add(a int, b int) int {
  return a + b
}

func main() {
	fmt.Println(add(42, 13))
  fmt.Println(add(42, 130))
}
