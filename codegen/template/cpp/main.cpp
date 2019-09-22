#include <fstream>
#include <iostream>
#include <codecvt>
#include <locale>

std::wstring_convert<std::codecvt_utf8<char32_t>, char32_t> cvt;
int main(int argc, char **argv)
{
    if (argc != 2) {
        return 0;
    }

    const char *test_file_path = argv[1];
    // Open the test file (contains UTF-8 encoded text)
    std::ifstream fs8(test_file_path);
    if (!fs8.is_open()) {
        return 0;
    }

    std::string line;
    while (getline(fs8, line)) {
        for (char32_t c: cvt.from_bytes(line)) // uses code point iterators
            std::printf("%#x ", c);
    }
}