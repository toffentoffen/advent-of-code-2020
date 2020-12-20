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
	// d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	var errors []int
	for _, t := range d.nearbyTickets {
		for _, fv := range t.fieldValues {
			anyValid := utils.Any(len(d.fields), func(i int) bool {
				return utils.Any(len(d.fields[i].rules), func(j int) bool {
					return d.fields[i].rules[j].from <= fv && d.fields[i].rules[j].to >= fv
				})
			})
			if !anyValid {
				errors = append(errors, fv)
			}
		}
	}
	fmt.Println(utils.IntSum(errors))
}

type rule struct {
	from int
	to   int
}
type field struct {
	name  string
	rules []rule
}
type ticket struct {
	fieldValues []int
}

type data struct {
	fields        []field
	yourTicket    ticket
	nearbyTickets []ticket
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
		line := s.Text()
		if line == "" {
			continue
		}
		if line == "your ticket:" {
			s.Scan()
			break
		}

		parts := strings.Split(line, ":")
		fieldName := parts[0]
		var rule1, rule2 rule
		rules := strings.Trim(parts[1], " ")
		utils.Sscanf(rules, "%d-%d or %d-%d", &rule1.from, &rule1.to, &rule2.from, &rule2.to)
		d.fields = append(d.fields, field{
			name:  fieldName,
			rules: []rule{rule1, rule2},
		})
	}
	d.yourTicket.fieldValues = utils.IntListWithSeparator(s.Text(), ",")

	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		if line == "nearby tickets:" {
			s.Scan()
		}
		d.nearbyTickets = append(d.nearbyTickets, ticket{fieldValues: utils.IntListWithSeparator(s.Text(), ",")})
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)
}
