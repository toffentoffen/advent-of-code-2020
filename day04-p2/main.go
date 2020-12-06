package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInvalidPassportsInputReader())
	//d, err := read(testValidPassportsInputReader())

	if err != nil {
		log.Fatal(err)
	}
	var valid int
	for _, p := range d.passports {
		if p.Valid() {
			valid++
		}
	}
	fmt.Printf("valid passwords: %d\n", valid)
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
				if !fieldValidators[vf](f) {
					return false
				}
				break
			}
		}
		if !found && vf != "cid" {
			return false
		}
	}
	return true
}

type fieldValidator func(string) bool

var byrValidatorFunc = func(s string) bool {
	var year int
	_, err := fmt.Sscanf(s, "byr:%d", &year)
	if err != nil {
		//log.Printf("could not extract byr from \"%s\": %v", s, err)
		return false
	}
	return year >= 1920 && year <= 2002
}

var iyrValidatorFunc = func(s string) bool {
	var year int
	_, err := fmt.Sscanf(s, "iyr:%d", &year)
	if err != nil {
		//log.Printf("could not extract iyr from \"%s\": %v", s, err)
		return false
	}
	return year >= 2010 && year <= 2020
}

var eyrValidatorFunc = func(s string) bool {
	var year int
	_, err := fmt.Sscanf(s, "eyr:%d", &year)
	if err != nil {
		//log.Printf("could not extract eyr from \"%s\": %v", s, err)
		return false
	}
	return year >= 2020 && year <= 2030
}

var hgtValidatorFunc = func(s string) bool {
	var height int
	var unit string
	_, err := fmt.Sscanf(s, "hgt:%d%s", &height, &unit)
	if err != nil {
		//log.Printf("could not extract hgt from \"%s\": %v", s, err)
		return false
	}
	if unit == "cm" {
		return height >= 150 && height <= 193
	}
	if unit == "in" {
		return height >= 59 && height <= 76
	}
	return false
}

var hclValidatorFunc = func(s string) bool {
	var color string
	_, err := fmt.Sscanf(s, "hcl:%s", &color)
	if err != nil {
		//log.Printf("could not extract hcl from \"%s\": %v", s, err)
		return false
	}
	r := regexp.MustCompile("#[0-9a-f]{6}")
	return r.MatchString(color)
}

var eclValidatorFunc = func(s string) bool {
	var eyeColor string
	_, err := fmt.Sscanf(s, "ecl:%s", &eyeColor)
	if err != nil {
		//log.Printf("could not extract ecl from \"%s\": %v", s, err)
		return false
	}

	switch eyeColor {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}

	return false
}

var pidValidatorFunc = func(s string) bool {
	var pid string
	_, err := fmt.Sscanf(s, "pid:%s", &pid)
	if err != nil {
		//log.Printf("could not extract pid from \"%s\": %v", s, err)
		return false
	}
	if len(pid) != 9 {
		return false
	}
	r := regexp.MustCompile("[0-9]{9}")
	return r.MatchString(pid)
}

var cidValidatorFunc = func(s string) bool {
	return true
}

var fieldValidators = map[string]fieldValidator{
	"byr": byrValidatorFunc,
	"iyr": iyrValidatorFunc,
	"eyr": eyrValidatorFunc,
	"hgt": hgtValidatorFunc,
	"hcl": hclValidatorFunc,
	"ecl": eclValidatorFunc,
	"pid": pidValidatorFunc,
	"cid": cidValidatorFunc,
}

type data struct {
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
		if s.Text() == "" {
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

var testInvalidPassportsInput = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`

var testValidPassportsInput = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

func testInvalidPassportsInputReader() io.Reader {
	return strings.NewReader(testInvalidPassportsInput)
}

func testValidPassportsInputReader() io.Reader {
	return strings.NewReader(testValidPassportsInput)
}
