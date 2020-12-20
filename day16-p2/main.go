package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	var errors []int
	var validTickets []ticket
	for _, t := range d.nearbyTickets {
		var validTicket = true
		for _, fv := range t.fieldValues {
			anyValid := utils.Any(len(d.fields), func(i int) bool {
				return utils.Any(len(d.fields[i].rules), func(j int) bool {
					return d.fields[i].rules[j].from <= fv && d.fields[i].rules[j].to >= fv
				})
			})
			if !anyValid {
				errors = append(errors, fv)
			}
			validTicket = validTicket && anyValid
		}
		if validTicket {
			validTickets = append(validTickets, t)
		}
	}

	for _, f := range d.fields {
		for i := 0; i < len(d.fields); i++ {
			all := utils.All(len(validTickets), func(ti int) bool {
				fv := validTickets[ti].fieldValues[i]
				return utils.Any(len(f.rules), func(j int) bool {
					return f.rules[j].from <= fv && f.rules[j].to >= fv
				})
			})
			if all {
				f.possibilities[i] = true
			}
		}
	}

	sort.Slice(d.fields, func(i, j int) bool {
		ci := utils.Count(len(d.fields[i].possibilities), func(p int) bool {
			return d.fields[i].possibilities[p]
		})
		cj := utils.Count(len(d.fields[j].possibilities), func(p int) bool {
			return d.fields[j].possibilities[p]
		})
		return ci < cj
	})
	var result = 1
	for i, f := range d.fields {
		c := utils.Count(len(f.possibilities), func(p int) bool {
			return f.possibilities[p]
		})
		if c == 1 {
			var pos int
			for i, p := range f.possibilities {
				if p {
					pos = i
				}
			}

			if strings.HasPrefix(f.name, "departure") {
				result *= d.yourTicket.fieldValues[pos]
			}
			for j := i + 1; j < len(d.fields); j++ {
				d.fields[j].possibilities[pos] = false
			}
		}
	}

	fmt.Println(result)
}

type rule struct {
	from int
	to   int
}
type field struct {
	name          string
	rules         []rule
	possibilities []bool
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
	for i := range d.fields {
		d.fields[i].possibilities = make([]bool, len(d.fields))
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`)
}
