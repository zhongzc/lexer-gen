//
// Created by zhong on 2019/9/23.
//

#ifndef CPPLEXERTEMPLATE_AUTOMATA_H
#define CPPLEXERTEMPLATE_AUTOMATA_H


#include <memory>
#include "CharReader.h"

class Automata {
public:
    virtual void runGreedy(std::shared_ptr<CharReader> reader) = 0;

    virtual std::string name() = 0;

    virtual ~Automata() = default;
};

class DOG : Automata {
    void runGreedy(std::shared_ptr<CharReader> reader) override;

    std::string name() override;
};

#endif //CPPLEXERTEMPLATE_AUTOMATA_H
