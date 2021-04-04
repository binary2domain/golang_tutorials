package main

import (
	"fmt"
)

func main() {
	// firstRank := 39
	// secondRank := 39

	// if firstRank < secondRank {
	// 	fmt.Println("\nFirst course is doing better than second course.")
	// } else if firstRank > secondRank {
	// 	fmt.Println("\nFirst course is doing bad.")
	// } else {
	// 	fmt.Println("\nBoth courses are doing equally bad.")
	// }

	// The firstRank and secondRank is scoped to this "if block"
	if firstRank, secondRank := 39, 694; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than second course.")
	} else if firstRank > secondRank {
		fmt.Println("\nFirst course is doing bad.")
	} else {
		fmt.Println("\nBoth courses are doing equally bad.")
	}
}
