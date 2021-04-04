package main

import (
	"fmt"
)

func main() {
	switch "docker" {
	case "linux":
		fmt.Println("\nHere are come recommended Linux courses ...")
	case "docker":
		fmt.Println("\nHere are come recommended Docker courses ...")
	case "windows":
		fmt.Println("\nHere are come recommended Windows courses ...")
	default:
		fmt.Println("Sorry could not find a match.")
	}
}
