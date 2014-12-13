#include "Parser.h"

#include <string>
#include <sstream>
#include <iostream>
#include <fstream>

using namespace std;

int main(int argc, char** argv){

	if (argc == 2){
		
		std::ifstream f;

		f.open(argv[1]);
		std::stringstream buffer;
		buffer << f.rdbuf();

		dom::Node *dom =  parse(buffer.str());

		f.close();

	}

	return 0;
}