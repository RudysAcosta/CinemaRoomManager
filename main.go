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

	seats := make([][]bool, rows)
	buildRoom(&seats, seatsForRow)

	totalSeats := rows * seatsForRow

	var totalIncome int

	if totalSeats > 60 {
		fistRows := math.Floor(float64(rows) / float64(2))
		lastRows := rows - int(fistRows)

		totalIncome = int(int(fistRows)*seatsForRow*10) + int(int(lastRows)*seatsForRow*8)

	} else {
		totalIncome = rows * seatsForRow * 10
	}

	// fmt.Println("Total income:")
	// fmt.Printf("$%d", totalIncome)

	fmt.Println()
	displaySeats(&seats, seatsForRow)
}

func buildRoom(seats *[][]bool, seatsForRow int) {
	for i := range *seats {
		(*seats)[i] = make([]bool, seatsForRow)
	}
}

func displaySeats(seats *[][]bool, seatsForRow int) {

	fmt.Printf(" ")
	for k := 1; k <= seatsForRow; k++ {
		fmt.Printf(" %d", k)
	}
	fmt.Println()

	for i := 0; i < len(*seats); i++ {
		fmt.Printf("%d", i+1)

		for j := 0; j < len((*seats)[i]); j++ {
			if (*seats)[i][j] {
				fmt.Printf(" B")
			} else {
				fmt.Printf(" S")
			}
		}
		fmt.Println()
	}
}
