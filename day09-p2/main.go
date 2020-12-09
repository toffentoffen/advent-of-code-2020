package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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

	encryption, err := d.encryption(weakness)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("encryption: %d\n", encryption)

}

type data struct {
	xmas []int
}

func (d data) weakness(preamble int) (int, error) {
	var weakness int
	for i := 0; i < len(d.xmas); i++ {
		weakness = d.xmas[i+preamble]
		hasSum := d.hasSum(d.xmas[i:i+preamble], weakness)
		if !hasSum {
			return weakness, nil
		}
	}
	return weakness, errors.New("could not find weakness")
}

func (d data) hasSum(ints []int, weakness int) bool {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints); j++ {
			if i != j {
				if ints[i]+ints[j] == weakness {
					return true
				}
			}
		}
	}
	return false
}

func (d data) encryption(weakness int) (int, error) {
	var encryption int
	for start := 0; start < len(d.xmas); start++ {
		for length := 1; length < len(d.xmas)-start; length++ {
			end := start + length
			if utils.IntSum(d.xmas[start:end]) == weakness {
				max := utils.Maximum(end-start, func(pos int) int {
					return d.xmas[pos+start]
				})
				min := utils.Minimum(end-start, func(pos int) int {
					return d.xmas[pos+start]
				})
				return max + min, nil
			}
		}
	}
	return encryption, errors.New("could not find encryption")
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
		d.xmas = append(d.xmas, utils.Atoi(s.Text()))
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
