//
// Created by zhong on 2019/9/23.
//

#ifndef CPPLEXERTEMPLATE_LEXER_H
#define CPPLEXERTEMPLATE_LEXER_H


#include <vector>
#include <memory>
#include "CharReader.h"
#include "Automata.h"

using TokenType = std::string;
using TokenValue = std::string;

class Token {
public:
    Token(TokenType type, TokenValue token);

public:
    TokenType type;
    TokenValue token;
};

class Lexer {
public:
    explicit Lexer(std::shared_ptr<CharReader> charReader);

    Token nextToken();

    bool hasNext();

    void skipWhitespace();

private:
    std::shared_ptr<CharReader> charReader;
    std::vector<std::unique_ptr<Automata>> automatas = std::vector<std::unique_ptr<Automata>>();
};


#endif //CPPLEXERTEMPLATE_LEXER_H
