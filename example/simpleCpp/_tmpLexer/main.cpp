#include <iostream>
#include <memory>
#include "CharReader.h"
#include "Lexer.h"

int main() {
    std::string s("if (a_for_apple == 10000) {\n"
                  "\tvar b_for_ball = 10086;\n"
                  "\treturn b_for_banana;\n"
                  "\t// a single-line comment\n"
                  "} else {\n"
                  "\t/*\n"
                  "\ta multi-line comment\n"
                  "\tcooooooool!\n"
                  "\t*/\n"
                  "\treturn 0;\n"
                  "}");
    auto cr = CharReader(s);
    auto lx = Lexer(cr);
    while (lx.hasNext()) {
        std::cout << lx.nextToken() << std::endl;
    }
}