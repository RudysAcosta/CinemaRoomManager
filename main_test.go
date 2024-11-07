package main

import (
	"testing"
)

func TestTickerBought_Add(t *testing.T) {
	ticker := TickerBought{}
	ticker.add(10)

	if ticker.Count != 1 {
		t.Errorf("Expected Count to be 1, got %d", ticker.Count)
	}

	if ticker.Income != 10 {
		t.Errorf("Expected Income to be 10, got %d", ticker.Income)
	}
}

func TestInitializeRoom(t *testing.T) {
	rows, seatsForRow := 3, 4
	seats := initializeRoom(rows, seatsForRow)

	if len(seats) != rows {
		t.Errorf("Expected %d rows, got %d", rows, len(seats))
	}

	for _, row := range seats {
		if len(row) != seatsForRow {
			t.Errorf("Expected %d seats per row, got %d", seatsForRow, len(row))
		}
	}
}

func TestBuyTicket(t *testing.T) {
	rows, seatsForRow := 5, 5
	seats := initializeRoom(rows, seatsForRow)
	row, seat := 2, 3
	seats[row-1][seat-1] = true // Mock seat as purchased

	if seats[row-1][seat-1] != true {
		t.Errorf("Expected seat %d in row %d to be true", seat, row)
	}

	// Attempt to buy the same seat
	buyTicket(seats, rows, seatsForRow)
	if !seats[row-1][seat-1] {
		t.Errorf("Seat %d in row %d should be marked as bought", seat, row)
	}
}

func TestCalculateSeatCost(t *testing.T) {
	// Test for more than 60 seats
	if cost := calculateSeatCost(70, 4, 8); cost != 8 {
		t.Errorf("Expected cost to be 8 for a second half seat, got %d", cost)
	}

	// Test for 60 or fewer seats
	if cost := calculateSeatCost(60, 2, 8); cost != 10 {
		t.Errorf("Expected cost to be 10 for seats <= 60, got %d", cost)
	}
}

func TestCalculateTotalIncome(t *testing.T) {
	// Case when total seats > 60
	if income := calculateTotalIncome(10, 10); income != 800 {
		t.Errorf("Expected total income to be 800, got %d", income)
	}

	// Case when total seats <= 60
	if income := calculateTotalIncome(5, 5); income != 250 {
		t.Errorf("Expected total income to be 250, got %d", income)
	}
}
