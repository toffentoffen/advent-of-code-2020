package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	maxID := utils.Maximum(len(d.seats), func(i int) int {
		return d.seats[i].id()
	})
	fmt.Printf("max seat id: %d\n", maxID)
}

type seat struct {
	row int8
	col int8
}

func (s seat) id() int {
	return int(s.row)*8 + int(s.col)
}

type data struct {
	seats []seat
}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()
	return read(f)
}

func read(r io.Reader) (*data, error) {
	var d data

	s := bufio.NewScanner(r)
	for s.Scan() {
		text := s.Text()
		seat := readSeat(text)
		d.seats = append(d.seats, *seat)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func readSeat(text string) *seat {
	text = strings.ReplaceAll(text, "F", "0")
	text = strings.ReplaceAll(text, "B", "1")
	text = strings.ReplaceAll(text, "R", "1")
	text = strings.ReplaceAll(text, "L", "0")
	rowN, err := strconv.ParseInt(text[0:7], 2, 8)
	if err != nil {
		log.Fatalf("could not parse seat %s: %v", text, err)
	}
	colN, err := strconv.ParseInt(text[7:10], 2, 4)
	if err != nil {
		log.Fatalf("could not parse seat %s: %v", text, err)
	}
	return &seat{
		row: int8(rowN),
		col: int8(colN),
	}
}
