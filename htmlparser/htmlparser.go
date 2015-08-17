package htmlparser

import (
	"fmt"
	"io"
	"text/scanner"
	"unicode"
)

type parser struct {
	input *scanner.Scanner
	next  rune
}

func scannerError(s *scanner.Scanner, msg string) {

	fmt.Println(msg)

}

func newParser(input io.Reader) *parser {
	var s scanner.Scanner
	s.Init(input)
	s.Error = scannerError
	s.Mode = scanner.ScanChars

	p := &parser{
		input: &s,
		next:  s.Scan(),
	}
	return p
}

type testRune func(rune) bool

func (p *parser) consumeWhile(test testRune) []rune {
	result := make([]rune, 0, 100)
	for p.next != scanner.EOF && test(p.next) {
		result = append(result, p.next)
		fmt.Printf("Parser Next(after):%q\n", p.next)
	}
	return result
}

func (p *parser) consumeWhitespace() {
	p.consumeWhile(unicode.IsSpace)
}

//func (p *parser) parseTagName() string {
//	return p.consumeWhile(func(r rune) bool {
//		return (unicode.IsLetter(r) || unicode.IsDigit(r))
//	})
//}
//
//func (p *parser) parseNode() *dom.Node {
//	switch p.Peek() {
//	case '<':
//		return p.parseElement()
//	default:
//		return p.parseText()
//	}
//}
//
//func (p *parser) parseText() *dom.Node {
//	dom.Text(p.consumeWhile(func(r rune) bool {
//		return (r != '<')
//	}))
//}
//
//func (p *Parse) parseElement() *dom.Node {
//	if p.next() != '<' {
//		//error on missing '<'
//	}
//
//	tagName := p.parseTagName()
//	attributes := p.parseAttributes()
//
//	if p.next() != '>' {
//		//error on missing '>'
//	}
//
//	children := p.parseNodes()
//
//	if p.next() != '<' {
//
//	}
//	if p.next() != '/' {
//
//	}
//	if p.parseTagName != tagName {
//
//	}
//	if p.next() != '>' {
//
//	}
//
//	return dom.Elem(tagName, attributes, children)
//}
//
//func (p *parser) parseAttribute() (string, string) {
//	name := p.parseTagName()
//
//	if p.next != '=' {
//
//	}
//
//	value := p.parseAttributeValue()
//
//	return name, value
//}
//
//func (p *parser) parseAttributeValue() string {
//	openQuote := p.next()
//	if openQuote != '"' || openQuote != '\'' {
//		//error here
//	}
//	value := p.consumeWhile(func(r rune) bool {
//		return r != openQuote
//	})
//	if p.next() != openQuote {
//		//error here
//	}
//	return value
//}
//
//func (p *parser) parseAttributes() map[string]string {
//	attributes = make(map[string]string)
//
//	for p.consumeWhitespace(); p.next() != '>'; {
//		name, value := p.parseAttribute()
//		attributes[name] = value
//	}
//
//	return attributes
//}
//
//func (p *parser) parseNodes() []*dom.Node {
//	nodes = make([]*dom.Node)
//	for p.consumeWhitespace; !p.end() || !p.startsWith("</"); {
//		append(nodes, p.parseNode)
//	}
//	return nodes
//}
//
//func Parse(source []byte) *dom.Node {
//	nodes := newParser(source).parseNodes()
//
//	if len(nodes) == 1 {
//		return nodes[0]
//	} else {
//		return dom.Elem("html", make(map[string]string), nodes)
//	}
//}
