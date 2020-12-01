package main

import "testing"

func TestFixExpenseReport(t *testing.T) {
	expenses := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	if FixExpenseReport(expenses, false) != 514579 {
		t.Errorf("got %v, want %v", FixExpenseReport(expenses, false), 514579)
	}

	if FixExpenseReport(expenses, true) != 241861950 {
		t.Errorf("got %v, want %v", FixExpenseReport(expenses, true), 241861950)
	}
}
