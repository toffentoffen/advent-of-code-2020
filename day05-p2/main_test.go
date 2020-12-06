package main

import "testing"

func TestReadSeat(t *testing.T) {
	s := readSeat("FBFBBFFRLR")
	if s.row != 44 {
		t.Errorf("unexpect seat row %d, was expecting 44", s.row)
	}
	if s.col != 5 {
		t.Errorf("unexpect seat col %d, was expecting 4", s.col)
	}
	if s.id() != 357 {
		t.Errorf("unexpect seat id %d, was expecting 357", s.id())
	}
}