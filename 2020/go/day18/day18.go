package day18

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

func Part1(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)
	var sum int
	for s.Scan() {
		sum += calculate(newParser(newLexer(s.Text()).generateToken()).parse1())
	}

	return strconv.Itoa(sum), nil
}

func Part2(f io.Reader) (string, error) {
	s := bufio.NewScanner(f)
	var sum int
	for s.Scan() {
		sum += calculate(newParser(newLexer(s.Text()).generateToken()).parse2())
	}

	return strconv.Itoa(sum), nil
}

//region Lexer
type tokenType int

const (
	Null tokenType = iota
	Number
	Plus
	Multiply
	LParen
	RParen
)

type token struct {
	typ tokenType
	v   int
}

type lexer struct {
	r *strings.Reader
	c rune
}

func newLexer(s string) *lexer {
	l := lexer{
		r: strings.NewReader(s),
	}
	l.advance()
	return &l
}

func (l *lexer) advance() {
	c, _, err := l.r.ReadRune()
	if err == io.EOF {
		l.c = 0
		return
	}
	aoc.NoError(err)
	l.c = c
	return
}

func (l *lexer) generateToken() (tokens []token) {
	for {
		switch {
		case unicode.IsSpace(l.c):
			// noop
		case unicode.IsDigit(l.c):
			var b bytes.Buffer
			for unicode.IsDigit(l.c) {
				b.WriteRune(l.c)
				l.advance()
			}
			v, err := strconv.Atoi(b.String())
			aoc.NoError(err)

			tokens = append(tokens, token{typ: Number, v: v})
			continue
		case l.c == '+':
			tokens = append(tokens, token{typ: Plus})
		case l.c == '*':
			tokens = append(tokens, token{typ: Multiply})
		case l.c == '(':
			tokens = append(tokens, token{typ: LParen})
		case l.c == ')':
			tokens = append(tokens, token{typ: RParen})
		case l.c == 0:
			return
		default:
			log.Fatalf("invalid token %c", l.c)
		}
		l.advance()

	}
	return
}

//endregion

//region Parser

type NumberNode int

func (n NumberNode) String() string {
	return strconv.Itoa(int(n))
}

type AddNode struct {
	a interface{}
	b interface{}
}

func (n AddNode) String() string {
	return fmt.Sprintf("(%v+%v)", n.a, n.b)
}

type MultiplyNode struct {
	a interface{}
	b interface{}
}

func (n MultiplyNode) String() string {
	return fmt.Sprintf("(%v*%v)", n.a, n.b)
}

type parser struct {
	tokens []token
	c      token
}

func newParser(t []token) *parser {
	return &parser{
		tokens: t,
	}
}

func (p *parser) advance() bool {
	if len(p.tokens) == 0 {
		p.c = token{typ: Null}
		return false
	}
	p.c = p.tokens[0]
	p.tokens = p.tokens[1:]
	return true
}

func (p *parser) parse1() interface{} {
	p.advance()
	result := p.expr1()
	if p.c.typ != Null {
		panic("invalid syntax")
	}
	return result
}

// expr1 evaluate expression where Plus and Multiplication does not have any precedence.
// That's why it's on the same level
func (p *parser) expr1() interface{} {
	result := p.factor1()

	for {
		switch p.c.typ {
		case Plus:
			p.advance()
			result = AddNode{a: result, b: p.factor1()}
		case Multiply:
			p.advance()
			result = MultiplyNode{a: result, b: p.factor1()}
		case RParen:
			return result
		case Null:
			fallthrough
		default:
			p.advance()
			return result
		}
	}
}

func (p *parser) factor1() interface{} {
	switch p.c.typ {
	case Number:
		result := NumberNode(p.c.v)
		p.advance()
		return result
	case LParen:
		p.advance()
		result := p.expr1() // recursive

		if p.c.typ != RParen {
			panic("expecting ')'")
		}
		p.advance()
		return result
	default:
		panic("factor invalid syntax")
	}
}

func (p *parser) parse2() interface{} {
	p.advance()
	result := p.expr2()
	if p.c.typ != Null {
		panic("invalid syntax")
	}
	return result
}

// expr2 evaluate expression where Multiplication has precedent, we add a new function term2 instead of
// adding it on the same variable
func (p *parser) expr2() interface{} {
	result := p.term2()

	for {
		switch p.c.typ {
		case Multiply:
			p.advance()
			result = MultiplyNode{a: result, b: p.term2()}
		case RParen:
			return result
		case Null:
			fallthrough
		default:
			p.advance()
			return result
		}
	}
}

func (p *parser) term2() interface{} {
	result := p.factor2()
	for {
		switch p.c.typ {
		case Plus:
			p.advance()
			result = AddNode{
				a: result,
				b: p.factor2(),
			}
		case Null:
			fallthrough
		default:
			return result
		}
	}
}

func (p *parser) factor2() interface{} {
	switch p.c.typ {
	case Number:
		result := NumberNode(p.c.v)
		p.advance()
		return result
	case LParen:
		p.advance()
		result := p.expr2()

		if p.c.typ != RParen {
			panic("expecting ')'")
		}
		p.advance()
		return result
	default:
		panic("factor invalid syntax")
	}
}

//endregion

//region interpreter

func calculate(node interface{}) int {
	switch n := node.(type) {
	case NumberNode:
		return int(n)
	case AddNode:
		return calculate(n.a) + calculate(n.b)
	case MultiplyNode:
		return calculate(n.a) * calculate(n.b)
	}
	panic("unknown node")
}

//endregion
