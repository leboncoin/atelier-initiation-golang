package main

import (
	"fmt"
)

func main() {
	fmt.Printf(fmt.Sprintf("%s\n", greet("World")))
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
