package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
var testMap = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func main() {
	//d, err := read(testMapReader())
	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	right := 3
	down := 1
	trees := 0
	column := 0
	for i := down; i<len(d.landMap); i=i+down {
		column = (column + right) % len(d.landMap[i])
		if d.landMap[i][column] == '#' {
			trees++
		}
	}
	fmt.Printf("trees: %d\n", trees)
}

type data struct{
	landMap []string
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
		d.landMap = append(d.landMap, strings.TrimSpace(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testMapReader() io.Reader {
	return strings.NewReader(testMap)
}