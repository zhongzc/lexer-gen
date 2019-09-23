//
// Created by zhong on 2019/9/23.
//

#include "Automata.h"

void DOG::runGreedy(std::shared_ptr<CharReader> reader) {
    auto old = reader->curIdx();

    char32_t c;
    s0:
    c = reader->nextChar();
    if (U'a' <= c && c <= U'b') {
        goto s1;
    }
    goto bad;

    s1:
    c = reader->nextChar();
    if (U'c' <= c && c <= U'd') {
        goto s2;
    }
    if (U'a' <= c && c <= U'b') {
        goto s1;
    }
    goto bad;

    s2:
    c = reader->nextChar();
    if (U'e' <= c && c <= U'f') {
        goto s1;
    }
    goto accept;

    bad:
    reader->setCurIdx(old);
    return;

    accept:
    return;
}

std::string DOG::name() {
    return "DOG";
}
