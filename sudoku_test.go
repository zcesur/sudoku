package main

import "testing"

func TestSolve(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		solution string
	}{
		// Test cases were retrieved from
		// “Valid Test Cases,” Valid Test Cases - Sudopedia Mirror. [Online].
		// Available: http://sudopedia.enjoysudoku.com/Valid_Test_Cases.html.
		// [Accessed: 15-Jan-2020]
		{
			"Completed Puzzle",
			"974236158638591742125487936316754289742918563589362417867125394253649871491873625",
			"974236158638591742125487936316754289742918563589362417867125394253649871491873625",
		},
		{
			"Last Empty Square",
			"2564891733746159829817234565932748617128.6549468591327635147298127958634849362715",
			"256489173374615982981723456593274861712836549468591327635147298127958634849362715",
		},
		{
			"Naked Singles",
			"3.542.81.4879.15.6.29.5637485.793.416132.8957.74.6528.2413.9.655.867.192.965124.8",
			"365427819487931526129856374852793641613248957974165283241389765538674192796512438",
		},
		{
			"Hidden Singles",
			"..2.3...8.....8....31.2.....6..5.27..1.....5.2.4.6..31....8.6.5.......13..531.4..",
			"672435198549178362831629547368951274917243856254867931193784625486592713725316489",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := ParseBoard(tt.input)
			if err != nil {
				t.Errorf("board could not be parsed: %v", err)
				return
			}

			bSolved, ok := b.Solve()
			if !ok {
				t.Errorf("solution not found")
				return
			}

			if bSolved.String() != tt.solution {
				t.Errorf("got %s, want %s", bSolved, tt.solution)
			}
		})
	}
}
