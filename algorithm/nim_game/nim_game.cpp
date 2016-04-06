#include <iostream>

class Solution {
public:
    bool canWinNim(int n) {
        return (n%4);
    }
};

int main(int argc, char **argv) {
	Solution s;
	std::cout<<s.canWinNim(4)<<std::endl;
	std::cout<<s.canWinNim(10)<<std::endl;
	std::cout<<s.canWinNim(11)<<std::endl;
	std::cout<<s.canWinNim(12)<<std::endl;
	std::cout<<s.canWinNim(13)<<std::endl;
	return 0;
}
