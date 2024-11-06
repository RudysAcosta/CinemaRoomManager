package main

import "fmt"

func main() {
	fmt.Println("Cinema:")
	fmt.Println("  1 2 3 4 5 6 7 8")

	for i := 1; i <= 7; i++ {
		fmt.Printf("%d ", i)
		for j := 1; j <= 8; j++ {
			fmt.Print("S ")
		}
		fmt.Println()
	}
}
