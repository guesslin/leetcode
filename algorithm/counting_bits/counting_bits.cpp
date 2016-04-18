#include <iostream>
#include <vector>

using std::vector;

class Solution {
public:
	vector<int> countBits(int num) {
		vector<int> r;
		r.push_back(0);
		if (num == 0) {
			return r;
		}
		for (int i = 0; i <= num; ++i) {
			r.push_back(r[i&(i-1)]+1);
		}
		return r;

	}
};

int main(int argc, char **argv) {
	Solution s;
	vector<int> result = s.countBits(123456);
	for (auto ite:result) {
		std::cout<<ite<<" ";
	}
	std::cout<<std::endl;
	return 0;
}

