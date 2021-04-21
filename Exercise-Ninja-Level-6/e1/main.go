package main

import "fmt"

func main() {
	x := foo()
	n, s := bar()
	fmt.Println(n, x, s)
}

func foo() int {
	return 2
}
func bar() (int, string) {
	return 2, "hello"
}
