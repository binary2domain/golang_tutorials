package main

// func testConnection(target string) (respTime float64 err error) {
// All functions must return error value
// If all goes well err should be returned as "nil"
//}

// if err != nil {
// 	handle the error properly
// }

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/home/santosh/text.txt")

	if err != nil {
		fmt.Println(err)
	}
}
