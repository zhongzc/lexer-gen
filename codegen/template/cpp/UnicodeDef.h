//
// Created by zhong on 2019/9/23.
//

#ifndef CPPLEXERTEMPLATE_UNICODEDEF_H
#define CPPLEXERTEMPLATE_UNICODEDEF_H

#include <locale>
#include <codecvt>

static std::wstring_convert<std::codecvt_utf8<char32_t>, char32_t> cvt32;

#endif //CPPLEXERTEMPLATE_UNICODEDEF_H
