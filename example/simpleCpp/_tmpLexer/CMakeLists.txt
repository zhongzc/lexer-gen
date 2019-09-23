cmake_minimum_required(VERSION 3.5)
project(cppLexerTemplate)

set(CMAKE_CXX_STANDARD 11)

add_executable(cppLexerTemplate main.cpp CharReader.cpp CharReader.h UnicodeDef.h Lexer.cpp Lexer.h Automata.cpp Automata.h)
