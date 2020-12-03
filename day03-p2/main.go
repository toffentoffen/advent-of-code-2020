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
	trees := 1

	for _, p := range []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		trees *= slopes(d, p.right, p.down)
	}

	fmt.Printf("trees: %d\n", trees)
}

func slopes(d *data, right int, down int) int {
	var column, trees int
	for i := down; i < len(d.landMap); i = i + down {
		column = (column + right) % len(d.landMap[i])
		if d.landMap[i][column] == '#' {
			trees++
		}
	}
	return trees
}

type data struct {
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
