#include <iostream>
#include <string>
#include <algorithm>

using std::string;

class Solution {
public:
    bool isPalindrome(string s) {
        std::transform(s.begin(), s.end(), s.begin(), std::tolower);
        for (int i = 0, j = s.size()-1; i <= j; i++, j--) {
            while(i < s.size() && !isAlphanumeric(s[i])) { i++; }
            while(j >= 0 && !isAlphanumeric(s[j])) { j--; }
            if(s[i] != s[j]) {
                return false;
            }
        }
        return true;
    }

    bool isAlphanumeric(char c) {
        if ('a' <= c && c <= 'z') {
            return true;
        }
        if ('0' <= c && c <= '9') {
            return true;
        }
        return false;
    }
};
