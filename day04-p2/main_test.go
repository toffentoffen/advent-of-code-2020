package main

import "testing"

func TestByrValidatorFunc(t *testing.T) {
	if !byrValidatorFunc("byr:2002") {
		t.Error("Expected \"byr:2002\" to be valid")
	}
	if byrValidatorFunc("byr:2003") {
		t.Error("Expected \"byr:2003\" to be invalid")
	}
}


func TestHgtValidatorFunc(t *testing.T) {
	if !hgtValidatorFunc("hgt:60in") {
		t.Error("Expected \"hgt:2002\" to be valid")
	}

	if !hgtValidatorFunc("hgt:190cm") {
		t.Error("Expected \"hgt:190cm\" to be valid")
	}

	if hgtValidatorFunc("hgt:190in") {
		t.Error("Expected \"hgt:190in\" to be invalid")
	}

	if hgtValidatorFunc("hgt:190") {
		t.Error("Expected \"hgt:190\" to be invalid")
	}
}

func TestHclValidatorFunc(t *testing.T) {
	if !hclValidatorFunc("hcl:#123abc") {
		t.Error("Expected \"hcl:#123abc\" to be valid")
	}

	if hclValidatorFunc("hcl:#123abz") {
		t.Error("Expected \"hcl:190in\" to be invalid")
	}

	if hclValidatorFunc("hcl:123abc") {
		t.Error("Expected \"hcl:123abc\" to be invalid")
	}
}

func TestEclValidatorFunc(t *testing.T) {
	if !eclValidatorFunc("ecl:brn") {
		t.Error("Expected \"ecl:brn\" to be valid")
	}
	if eclValidatorFunc("ecl:wat") {
		t.Error("Expected \"byr:wat\" to be invalid")
	}
}

func TestPidValidatorFunc(t *testing.T) {
	if !pidValidatorFunc("pid:000000001") {
		t.Error("Expected \"pid:000000001\" to be valid")
	}
	if pidValidatorFunc("pid:0123456789") {
		t.Error("Expected \"byr:0123456789\" to be invalid")
	}
	if pidValidatorFunc("pid:0000001") {
		t.Error("Expected \"byr:0000001\" to be invalid")
	}
}