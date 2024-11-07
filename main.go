package main

import (
	"fmt"
	"math"
	"os"
)

type TickerBought struct {
	Count  int
	Income int
}

func (t *TickerBought) add(cost int) {
	t.Count = t.Count + 1
	t.Income = t.Income + cost
}

func main() {

	var rows, seatsForRow int

	tickerBought := TickerBought{
		Count:  0,
		Income: 0,
	}

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
			costSeat := getTicket(&seats, rows, seatsForRow)

			tickerBought.add(costSeat)

			fmt.Println()
			fmt.Printf("Ticket price: $%d\n", costSeat)

		} else if option == 3 {
			statistics(&tickerBought, rows, seatsForRow)
		} else if option == 0 {
			os.Exit(0)
		}
	}
}

func menu() {
	fmt.Println()
	fmt.Println("1. Show the seats")
	fmt.Println("2. Buy a ticket")
	fmt.Println("3. Statistics")
	fmt.Println("0. Exit")
}

func getTicket(seats *[][]bool, rows, seatsForRow int) int {
	var row, seat int

	for {
		fmt.Println("Enter a row number:")
		fmt.Scanln(&row)

		fmt.Println("Enter a seat number in that row:")
		fmt.Scanln(&seat)

		if row > rows || seat > seatsForRow {
			fmt.Printf("\nWrong input!\n\n")
			continue
		}

		if !(*seats)[row-1][seat-1] {
			(*seats)[row-1][seat-1] = true
			break
		}

		fmt.Printf("\nThat ticket has already been purchased!\n\n")
	}

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

func statistics(ticker *TickerBought, rows, seatsForRow int) {

	totalIncome := totalIncome(rows, seatsForRow)
	totalSeats := rows * seatsForRow

	percentage := (float64(ticker.Count) / float64(totalSeats)) * 100

	fmt.Println()
	fmt.Printf("Number of purchased tickets: %d\n", ticker.Count)
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Printf("Current income: $%d\n", ticker.Income)
	fmt.Printf("Total income: $%d\n", totalIncome)
}

func totalIncome(rows, seatsForRow int) int {
	if rows*seatsForRow >= 60 {
		fistRows := math.Floor(float64(rows) / float64(2))
		lastRows := rows - int(fistRows)
		return int(int(fistRows)*seatsForRow*10) + int(int(lastRows)*seatsForRow*8)
	} else {
		return rows * seatsForRow * 10
	}
}

// Retona el numero del asiento
func getNumSeat(row, seatsForRow, seat int) int {
	seatNumber := (row - 1*seatsForRow) + seat
	return seatNumber
}
