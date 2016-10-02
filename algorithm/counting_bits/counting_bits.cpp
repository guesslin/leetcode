#include <iostream>
#include <vector>

using std::vector;
using std::cout;
using std::endl;

class Solution {
public:
	vector<int> countBits(int num) {
		vector<int> r(num+1, 0);
		if (num == 0) {
			return r;
		}
		for (int i = 1; i <= num; ++i) {
			r[i] = r[i&(i-1)]+1;
		}
		return r;

	}
};

int main(void) {
	Solution s;
	vector<int> result = s.countBits(123456);
	for (auto ite:result) {
		std::cout<<ite<<" ";
	}
	cout << endl;
	return 0;
}

