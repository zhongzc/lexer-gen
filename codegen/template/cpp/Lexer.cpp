//
// Created by zhong on 2019/9/23.
//

#include "Lexer.h"

#include <utility>

Lexer::Lexer(std::shared_ptr<CharReader> charReader) : charReader(std::move(charReader)) {
    this->automatas.push_back(std::make_unique<DOG>());

}

Token Lexer::nextToken() {
    if (!this->hasNext()) {
        throw std::runtime_error("reach EOF");
    }
    this->skipWhitespace();

    for (const auto &a : this->automatas) {
        auto o = this->charReader->curIdx();
        a->runGreedy(this->charReader);
        auto n = this->charReader->curIdx();

        if (o == n) {
            continue;
        } else {
            auto res = Token(a->name(), CharReader::subString(o, n));
            this->skipWhitespace();
            return res;
        }
    }

    throw std::runtime_error("unexpected token");
}

bool Lexer::hasNext() {
    return this->charReader->peek() != 0;
}

void Lexer::skipWhitespace() {
    auto p = this->charReader->peek();
    while (std::iswspace(p)) {
        this->charReader->nextChar();
        p = this->charReader->peek();
    }
}


Token::Token(TokenType type, TokenValue token) : type(std::move(type)), token(std::move(token)) {}

