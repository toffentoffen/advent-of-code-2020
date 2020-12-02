package main

import (
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
	pos1     int
	pos2     int
	letter   rune
	password string
}

func (dp databasePasswords) valid() bool {
	var occurences int
	if rune(dp.password[dp.pos1-1]) == dp.letter {
		occurences++
	}
	if rune(dp.password[dp.pos2-1]) == dp.letter {
		occurences++
	}
	return occurences == 1
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
		var pos1, pos2 int
		var c rune
		var password string
		_, err := fmt.Sscanf(s.Text(), "%d-%d %c: %s", &pos1, &pos2, &c, &password)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		cps := databasePasswords{
			pos1:     pos1,
			pos2:     pos2,
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
