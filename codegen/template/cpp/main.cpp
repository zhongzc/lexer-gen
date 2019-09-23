#include <iostream>
#include <memory>
#include "CharReader.h"

int main() {
    std::string s("ac ac ac ac");
    auto cr = std::make_shared<CharReader>(s);

}