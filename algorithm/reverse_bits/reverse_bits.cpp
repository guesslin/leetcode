#include <iostream>
#include <cstdint>

class Solution {
	public:
		uint32_t reverseBits(uint32_t n) {
			uint32_t r = 0;
			for (int i = 0; i < 32; i++) {
				r = r<<1;
				r+=n&1;
				n = n>>1;
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
