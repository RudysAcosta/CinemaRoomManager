package main

import (
	"fmt"
	"math"
)

func main() {

	var rows, seatsForRow int

	fmt.Println("Enter the number of rows:")
	fmt.Scanln(&rows)

	fmt.Println("Enter the number of seats in each row:")
	fmt.Scanln(&seatsForRow)

	seats := rows * seatsForRow
	var totalIncome int

	if seats > 60 {
		fistRows := math.Floor(float64(rows) / float64(2))
		lastRows := rows - int(fistRows)

		totalIncome = int(int(fistRows)*seatsForRow*10) + int(int(lastRows)*seatsForRow*8)

	} else {
		totalIncome = rows * seatsForRow * 10
	}

	fmt.Println("Total income:")
	fmt.Printf("$%d", totalIncome)
}
