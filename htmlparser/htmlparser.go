package htmlparser

import (
	"unicode/utf8"
	"unicode"
	"strings"
	"fmt"
	"toy-browser-engine/dom"
)

type Parser struct {
	Input    *string
	Position uint
}


func newParser(input *string) *Parser {
	return &Parser{
		Position: 0,
		Input:    input,
	}
}

func (p *Parser) peek() rune {
	r, size := utf8.DecodeRuneInString(p.Input[p.Position:])
	return r
}

func (p *Parser) startsWith(str *string) bool {
	return strings.HasPrefix(p.Input[p.Position:], str)
}

func (p *Parser) end() bool {
	return (p.Position >= len(p.Input)
}

func (p *Parser) next() rune {
	r, size := utf8.DecodeRuneInString(p.Input[p.Position:])
	p.Position = p.Position + uint(size)
	return r	
}

type testRune func(rune) bool
	
func (p *Parser) consumeWhile(test testRune) string {
	tempString := make([]rune, 50, 100)
	for !p.End() && test(p.Peek()) {
		append(tempString, p.Next())
	}
	return fmt.Sprintf("%c", tempString)
}

func (p *Parser) consumeWhitespace() {
	p.consumeWhile(unicode.IsSpace)
}	 

func (p *Parser) parseTagName() string {
	return p.consumeWhile(func(r rune) bool {
		return (unicode.IsLetter(r) || unicode.IsDigit(r))
	})
}

func (p *Parser) parseNode() *dom.Node {
	switch p.Peek() {
		case '<':
			return p.parseElement()
		default:
			return p.parseText()
	}
}

func (p *Parser) parseText() *dom.Node {
	dom.Text(p.consumeWhile(func(r rune) bool {
		return (r != '<')
	}))
}

func (p *Parse) parseElement() *dom.Node {
	  if p.next() != '<' {
	  //error on missing '<'
	  }
	  
	  tagName := p.parseTagName()
	  attributes := p.parseAttributes()
	  
	  if p.next() != '>' {
	  //error on missing '>'
	  }
	  
	  children := p.parseNodes()
	  
	  if p.next() != '<' {
		  
	  }
	  if p.next() != '/'{
		  
	  }
	  if p.parseTagName != tagName {
		  
	  }
	  if p.next() != '>' {
		  
	  }
	  
	  return dom.Elem(tagName, attributes, children)
}

func (p *Parser) parseAttribute() (string, string) {
	name := p.parseTagName();
	
	if p.next != '=' {
		
	}
	
	value := p.parseAttributeValue()
	
	return (name, value)
}

func (p *Parser) parseAttributeValue() string {
	openQuote := p.next()
	if (openQuote != '"' || openQuote != '\'') {
		//error here
	}
	value := p.consumeWhile(func(r rune) bool {
		return r != openQuote
	})
	if p.next() != openQuote {
		//error here
	}
	return value
}

func (p *Parser) parseAttributes() map[string]string {
	attributes = make(map[string]string)
	
	for p.consumeWhitespace(); p.next() != '>' {
		name, value := p.parseAttribute();
		attributes[name] = value
	}
	
	return attributes
}

func (p *Parser) parseNodes []*dom.Node {
	nodes = make([]*dom.Node)
	for p.consumeWhitespace; !p.end() || !p.startsWith("</") {
		append(nodes, p.parseNode)
	}
	return nodes
}

func Parse(source string) *dom::Node {
	nodes := NewParser(source).parseNodes()
	
	if len(nodes) == 1 {
		
	}
}