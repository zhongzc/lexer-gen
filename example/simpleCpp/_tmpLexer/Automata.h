// generated by cppGen.go - do not edit
#ifndef CPPLEXERTEMPLATE_AUTOMATA_H
#define CPPLEXERTEMPLATE_AUTOMATA_H


#include <memory>
#include "CharReader.h"

class Automata {
public:
    virtual void runGreedy(CharReader& reader) = 0;

    virtual std::string name() = 0;

    virtual ~Automata() = default;
};


class IF : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class ELSE : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class LPAREN : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class RPAREN : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class LBRACKET : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class RBRACKET : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class SEMICOLON : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class VAR : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class RETURN : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class EQ : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class ASSIGN : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class NUM : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class IDENT : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class MCOMMENT : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};
class SCOMMENT : public Automata {
    void runGreedy(CharReader& reader) override;

    std::string name() override;
};

#endif //CPPLEXERTEMPLATE_AUTOMATA_H