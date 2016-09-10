#include <iostream>
#include <string>

using std::string;

class Solution {
public:
    char findTheDifference(string s, string t) {
        char res = 0;
        for(auto ite:s) {
            res ^= ite;
        }
        for(auto ite:t) {
            res ^= ite;
        }
        return res;
    }
};
