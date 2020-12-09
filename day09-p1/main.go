package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	weakness, err := d.weakness(25)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("weakness: %d\n", weakness)
}

type data struct{
	xmas []int64
}

func (d data) weakness(preamble int) (int64, error) {
	var weakness int64
	for i := 0; i < len(d.xmas); i++ {
		weakness = d.xmas[i+preamble]
		hasSum := d.hasSum(d.xmas[i:i+preamble], weakness)
		if ! hasSum {
			return weakness, nil
		}
	}
	return weakness, errors.New("could not find weakness")
}

func (d data) hasSum(ints []int64, weakness int64) bool {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints); j++ {
			if i != j {
				if ints[i] + ints[j] == weakness {
					return true
				}
			}
		}
	}
	return false
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
		val, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatalf("could not parse %s: %v", s.Text(), err)
		}
		d.xmas = append(d.xmas, val)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)
}