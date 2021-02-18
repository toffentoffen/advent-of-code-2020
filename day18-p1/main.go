package main

import (
	"advent-of-code-2020/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	d, err := readInput("input.txt")
	//d, err := read(testInputReader())
	if err != nil {
		log.Fatal(err)
	}
	var sum int
	for _, e := range d.rawExpressions {
		er := newExpressionLexer(strings.NewReader(e))
		tokens := er.tokens()

		exp := parseExp(&tokens)
		value := exp.evaluate()
		fmt.Printf("%s = %d\n", e, value)
		sum += value
	}
	fmt.Printf("Sum: %d\n", sum)
}

type expressionLexer struct {
	reader *bufio.Reader
}

func newExpressionLexer(reader io.Reader) *expressionLexer {
	return &expressionLexer{
		reader: bufio.NewReader(reader),
	}
}

type tokenType int

const (
	EOF = iota
	ILLEGAL
	INT

	// Infix ops
	ADD // +
	MUL // *

	LPAR // (
	RPAR // )
)

func (tt tokenType) in(tts ...tokenType) bool {
	for _, t := range tts {
		if t == tt {
			return true
		}
	}
	return false
}

type token struct {
	tokenType
	value string
}

func (er *expressionLexer) lex() (tokenType, string) {
	for {
		r, _, err := er.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return EOF, ""
			}

			// at this point there isn't much we can do, and the compiler
			// should just return the raw error to the user
			panic(err)
		}
		switch r {
		case '+':
			return ADD, "+"
		case '*':
			return MUL, "*"
		case '(':
			return LPAR, "("
		case ')':
			return RPAR, ")"
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				er.backup()
				return INT, er.lexInt()
			}
		}
	}
}

func (er *expressionLexer) lexInt() string {
	var value string
	for {
		r, _, err := er.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the value
				return value
			}
		}

		if unicode.IsDigit(r) {
			value = value + string(r)
		} else {
			// scanned something not in the integer
			er.backup()
			return value
		}
	}
}

func (er *expressionLexer) backup() {
	if err := er.reader.UnreadRune(); err != nil {
		panic(err)
	}
}

func (er *expressionLexer) tokens() []token {
	var tokens []token
	for {
		t, v := er.lex()
		tokens = append(tokens, token{
			tokenType: t,
			value:     v,
		})
		if t == EOF {
			break
		}
	}

	return tokens
}

type expression interface {
	evaluate() int
}

type valueExpression struct {
	value int
}

func (e valueExpression) evaluate() int {
	return e.value
}

type sumExpression struct {
	lExp expression
	rExp expression
}

func (s sumExpression) evaluate() int {
	return sumExpressionFunc(s.lExp, s.rExp)
}

type multiplyExpression struct {
	lExp expression
	rExp expression
}

func (s multiplyExpression) evaluate() int {
	return multiplyExpressionFunc(s.lExp, s.rExp)
}

var sumExpressionFunc = func(r expression, l expression) int {
	re := r.evaluate()
	le := l.evaluate()
	return re + le
}

var multiplyExpressionFunc = func(rExp expression, lExp expression) int {
	re := rExp.evaluate()
	le := lExp.evaluate()
	return re * le
}

func parseExp(tokens *[]token) expression {
	var lExp = parseExpOperation(tokens)

	for (*tokens)[0].tokenType.in(ADD, MUL){
		t := (*tokens)[0]
		*tokens = (*tokens)[1:]
		var exp expression
		switch t.tokenType {
		case ADD:
			exp = sumExpression{
				lExp: lExp,
				rExp: parseExpOperation(tokens),
			}
		case MUL:
			exp = multiplyExpression{
				lExp: lExp,
				rExp: parseExpOperation(tokens),
			}
		default:
			panic(fmt.Sprintf("invalid token %s", t.value))
		}
		lExp = exp
	}
	return lExp
}

func parseExpOperation(tokens *[]token) expression {
	t := (*tokens)[0]
	*tokens = (*tokens)[1:]
	if t.tokenType == INT {
		return valueExpression{value: utils.Atoi(t.value)}
	}
	if t.tokenType != LPAR {
		panic(fmt.Sprintf("invalid token %s", t.value))
	}
	exp := parseExp(tokens)

	t = (*tokens)[0]
	*tokens = (*tokens)[1:]
	if t.tokenType != RPAR {
		panic(fmt.Sprintf("invalid token %s", t.value))
	}
	return exp
}

type data struct {
	rawExpressions []string
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
		d.rawExpressions = append(d.rawExpressions, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return &d, s.Err()
}

func testInputReader() io.Reader {
	return strings.NewReader(`1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`)
}
