#include <iostream>
#include <cstdint>


unsigned int bits[] = {0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15};

class Solution {
	public:
		uint32_t reverseBits(uint32_t n) {
			uint32_t r = 0;
			for (int i = 0; i < 8; i++) {
				r = r<<4;
				r+=bits[n&15];
				n = n>>4;
			}
			return r;
		}
};

int main(int argc, const char *argv[])
{
	Solution s;
	std::cout<<s.reverseBits(43261596)<<std::endl;
	return 0;
}
