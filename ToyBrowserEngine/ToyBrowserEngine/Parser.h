#ifndef _PARSER_H_
#define _PARSER_H_

#include "Node.h"

#include <string>
#include <sstream>


dom::Node *parse(std::string source);

class Parser {
private:
	std::stringstream input;


public: 

	Parser(std::string input_) : input(input_) {}

	dom::Node *parseNode();

	dom::Text *parseText();

	dom::Element *parseElement();

	dom::attr parseAttr();

	std::string parseAttrValue();

	dom::AttrMap parseAttributes();

	std::vector<dom::Node*> parseNodes();

	dom::Node parse(std::string source);

protected:

	char nextChar();

	bool startsWith(std::string str);

	bool eof();

	char consumeChar();

	std::string consumeWhile(bool(*f)(char));
	
	void consumeWhitespace();
	
	std::string parseTagName();
};


#endif /* _PARSER_H_*/