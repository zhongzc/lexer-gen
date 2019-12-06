// generated by cppGen.go - do not edit
#ifndef CPPLEXERTEMPLATE_CHARREADER_H
#define CPPLEXERTEMPLATE_CHARREADER_H

#include "UnicodeDef.h"

class CharReader {
public:
    CharReader(const std::u32string::iterator &begin,
               const std::u32string::iterator &end);
    explicit CharReader(const std::string &s);

    char32_t peek();

    char32_t nextChar();

    void setCurIdx(const std::u32string::iterator &c);

    static std::string subString(const std::u32string::iterator &b, const std::u32string::iterator &e);

    const std::u32string::iterator &curIdx();

private:
    std::u32string str;
    std::u32string::iterator cur;
    std::u32string::iterator end;
};


#endif //CPPLEXERTEMPLATE_CHARREADER_H