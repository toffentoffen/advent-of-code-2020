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
	for _, x := range d.ints {
		for _, y := range d.ints {
			for _, z := range d.ints {
				if x+y+z == 2020 {
					fmt.Printf("result: %d\n", x*y*z)
					return
				}
			}
		}
	}
}

type data struct {
	ints []int
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
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		d.ints = append(d.ints, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}
