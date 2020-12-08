package main

import (
	"advent-of-code-2020/utils"
	"bufio"
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
	fmt.Printf("read %d bags\n", len(d.bags))
	fmt.Printf("colors: %d\n", len(solve(d, "shiny gold")))
}

type bag struct {
	color     string
	contains map[string]int // key is bag color value is quantity
}
type data struct {
	bags map[string]bag
}

func solve(d *data, mustContain string) map[string]bool {
	var bagColors = make(map[string]bool)

	for color, b := range d.bags {
		if color != mustContain {
			if _, ok := b.contains[mustContain]; ok {
				bagColors[color] =true
				colors := solve(d, color)
				for c, _ := range colors {
					bagColors[c] = true
				}
			}
		}
	}
	return bagColors
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
	var d = data{
		map[string]bag{},
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		b := readBag(s.Text())
		d.bags[b.color] = *b
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func readBag(s string) *bag {
	fields := strings.Fields(s)
	var bg = bag{
		color:     strings.Join(fields[0:2], " "),
		contains: map[string]int{},
	}
	var i int = 4
	for i < len(fields) {
		if fields[i] == "no" {
			break
		}
		color := strings.Join(fields[i+1:i+3], " ")
		bg.contains[color] = utils.Atoi(fields[i])
		i += 4
	}
	return &bg
}

func testInputReader() io.Reader {
	return strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)
}
