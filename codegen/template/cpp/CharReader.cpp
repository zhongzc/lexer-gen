//
// Created by zhong on 2019/9/23.
//

#include "CharReader.h"

CharReader::CharReader(
        const std::u32string::iterator &begin,
        const std::u32string::iterator &end) : cur(begin), end(end) {}

CharReader::CharReader(const std::string &str) {
    auto us = cvt32.from_bytes(str);
    this->cur = us.begin();
    this->end = us.end();
}

char32_t CharReader::peek() {
    if (this->cur == this->end) {
        return 0;
    }
    return (*(this->cur + 1));
}

char32_t CharReader::nextChar() {
    if (this->cur == this->end) {
        return 0;
    }
    auto c = *this->cur;
    this->cur += 1;
    return c;
}

void CharReader::setCurIdx(const std::u32string::iterator &c) {
    this->cur = c;
}

std::string CharReader::subString(const std::u32string::iterator &b, const std::u32string::iterator &e) {
    return cvt32.to_bytes(&(*(b)), &(*(e)));
}

const std::u32string::iterator &CharReader::curIdx() {
    return this->cur;
}



