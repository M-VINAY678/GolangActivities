package main

import "testing"

func TestDeposit(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 100, 100},
		{"negative", -100, -100},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			d := &Account{Balance: 0}
			d.Deposit(tt.input)
			if d.Balance != tt.expected {
				t.Errorf("FunctionName(%v) = %v; want %v", tt.input, d.Balance, tt.expected)
			}
		})
	}
}
