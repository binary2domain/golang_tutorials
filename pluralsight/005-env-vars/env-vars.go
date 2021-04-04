package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Environ())
	for a, env := range os.Environ() {
		fmt.Println(a, env)
	}

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USER")
	fmt.Println("Username is", name)
}
