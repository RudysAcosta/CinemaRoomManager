package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var rows, seatsForRow int

	fmt.Println("Enter the number of rows:")
	fmt.Scanln(&rows)

	fmt.Println("Enter the number of seats in each row:")
	fmt.Scanln(&seatsForRow)

	seats := make([][]bool, rows)
	buildRoom(&seats, seatsForRow)

	for {
		var option int
		menu()
		fmt.Scanln(&option)

		if option == 1 {
			displaySeats(&seats, seatsForRow)
		} else if option == 2 {
			fmt.Println()
			costSeat := getTicket(&seats)

			fmt.Println()
			fmt.Printf("Ticket price: $%d\n", costSeat)

		} else if option == 0 {
			os.Exit(0)
		}
	}
}

func menu() {
	fmt.Println()
	fmt.Println("1. Show the seats")
	fmt.Println("2. Buy a ticket")
	fmt.Println("0. Exit")
}

func getTicket(seats *[][]bool) int {
	var row, seat int

	fmt.Println("Enter a row number:")
	fmt.Scanln(&row)

	fmt.Println("Enter a seat number in that row:")
	fmt.Scanln(&seat)

	(*seats)[row-1][seat-1] = true

	return getSeatCost(seats, row)
}

func getSeatCost(seats *[][]bool, row int) int {

	totalSeats := len(*seats) * len((*seats)[0])

	if totalSeats > 60 {
		fistRows := math.Floor(float64(len(*seats)) / float64(2))
		if row > int(fistRows) {
			return 8
		}
	}

	return 10
}

func buildRoom(seats *[][]bool, seatsForRow int) {
	for i := range *seats {
		(*seats)[i] = make([]bool, seatsForRow)
	}
}

func displaySeats(seats *[][]bool, seatsForRow int) {
	fmt.Println()
	fmt.Println("Cinema:")
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
