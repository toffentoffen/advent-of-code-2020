package main

import (
	"../utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var n int
	for _, dp := range d.passwords {
		if dp.valid() {
			n++
		}
	}

	fmt.Printf("valid passwords: %d\n", n)
}

type databasePasswords struct {
	min      int
	max      int
	letter   rune
	password string
}

func (dp databasePasswords) valid() bool {
	cc := utils.CharCounts(dp.password)

	if cc[dp.letter] >= dp.min && cc[dp.letter] <= dp.max {
		return true
	}
	return false
}

type data struct {
	passwords []databasePasswords
}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data

	s := bufio.NewScanner(f)
	for s.Scan() {
		var min, max int
		var c rune
		var password string
		_, err := fmt.Sscanf(s.Text(), "%d-%d %c: %s", &min, &max, &c, &password)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		cps := databasePasswords{
			min:      min,
			max:      max,
			letter:   c,
			password: password,
		}

		d.passwords = append(d.passwords, cps)

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}
