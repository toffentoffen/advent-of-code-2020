package main

import (
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
	var valids int
	for _, p := range d.passports {
		if p.Valid() {
			valids++
		}
	}
	fmt.Printf("valid passwords: %d\n", valids)
}

type passport struct {
	fields []string
}
var validFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
func (p passport) Valid() bool {
	if len(p.fields) < len(validFields)-1 {
		return false
	}
	for _, vf := range validFields {
		found := false
		for _, f := range p.fields {
			if strings.HasPrefix(f, vf) {
				found = true
				break
			}
		}
		if !found && vf != "cid" {
			return false
		}
	}
	return true
}
type data struct{
	passports []passport
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
	var fields []string
	for s.Scan() {
		if s.Text() == ""{
			d.passports = append(d.passports, passport{fields: fields})
			fields = []string{}
			continue
		}
		fields = append(fields, strings.Fields(s.Text())...)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	d.passports = append(d.passports, passport{fields: fields})
	return &d, s.Err()
}

var testInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func testInputReader() io.Reader {
	return strings.NewReader(testInput)
}