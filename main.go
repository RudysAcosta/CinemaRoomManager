package main

import (
	"fmt"
	"os"
)

type TickerBought struct {
	Count  int
	Income int
}

func (t *TickerBought) add(cost int) {
	t.Count++
	t.Income += cost
}

func main() {
	var rows, seatsForRow int
	fmt.Println("Enter the number of rows:")
	fmt.Scanln(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scanln(&seatsForRow)

	seats := initializeRoom(rows, seatsForRow)
	tickerBought := TickerBought{}

	for {
		switch option := displayMenu(); option {
		case 1:
			displaySeats(seats)
		case 2:
			costSeat := buyTicket(seats, rows, seatsForRow)
			tickerBought.add(costSeat)
			fmt.Printf("Ticket price: $%d\n", costSeat)
		case 3:
			displayStatistics(tickerBought, rows, seatsForRow)
		case 0:
			os.Exit(0)
		}
	}
}

func displayMenu() int {
	fmt.Println("\n1. Show the seats")
	fmt.Println("2. Buy a ticket")
	fmt.Println("3. Statistics")
	fmt.Println("0. Exit")
	var option int
	fmt.Scanln(&option)
	return option
}

func initializeRoom(rows, seatsForRow int) [][]bool {
	seats := make([][]bool, rows)
	for i := range seats {
		seats[i] = make([]bool, seatsForRow)
	}
	return seats
}

func buyTicket(seats [][]bool, rows, seatsForRow int) int {
	var row, seat int
	for {
		fmt.Println("Enter a row number:")
		fmt.Scanln(&row)
		fmt.Println("Enter a seat number in that row:")
		fmt.Scanln(&seat)
		if validSeat(row, seat, rows, seatsForRow) && !seats[row-1][seat-1] {
			seats[row-1][seat-1] = true
			break
		}
		fmt.Println("Invalid input or seat already purchased, please try again.")
	}
	return calculateSeatCost(rows*seatsForRow, row, rows)
}

func validSeat(row, seat, rows, seatsForRow int) bool {
	return row > 0 && row <= rows && seat > 0 && seat <= seatsForRow
}

func calculateSeatCost(totalSeats, row, rows int) int {
	if totalSeats > 60 && row > rows/2 {
		return 8
	}
	return 10
}

func displaySeats(seats [][]bool) {
	fmt.Println("\nCinema:")
	fmt.Print("  ")
	for i := 1; i <= len(seats[0]); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for i, row := range seats {
		fmt.Printf("%d ", i+1)
		for _, seat := range row {
			if seat {
				fmt.Print("B ")
			} else {
				fmt.Print("S ")
			}
		}
		fmt.Println()
	}
}

func displayStatistics(ticker TickerBought, rows, seatsForRow int) {
	totalSeats := rows * seatsForRow
	totalIncome := calculateTotalIncome(rows, seatsForRow)
	percentage := (float64(ticker.Count) / float64(totalSeats)) * 100

	fmt.Printf("\nNumber of purchased tickets: %d\n", ticker.Count)
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Printf("Current income: $%d\n", ticker.Income)
	fmt.Printf("Total income: $%d\n", totalIncome)
}

func calculateTotalIncome(rows, seatsForRow int) int {
	if rows*seatsForRow > 60 {
		firstRows := rows / 2
		lastRows := rows - firstRows
		return firstRows*seatsForRow*10 + lastRows*seatsForRow*8
	}
	return rows * seatsForRow * 10
}
