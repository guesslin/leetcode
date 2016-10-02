#include <iostream>
#include <vector>
#include <algorithm>

using std::vector;
using std::max;
using std::cout;
using std::endl;

class Solution {
public:
	int maxProfit(vector<int>& prices) {
		if (prices.size() == 0) {
			return 0
		}
		vector<int> s0(prices.size(), 0);
		vector<int> s1(prices.size(), -prices[0]);
		vector<int> s2(prices.size(), INT_MIN);
		for (int i = 1; i < prices.size(); ++i) {
			s0[i] = max(s0[i-1], s2[i-1]);
			s1[i] = max(s1[i-1], s0[i-1]-prices[i]);
			s2[i] = s1[i-1]+prices[i];
		}
		return max(s2[prices.size()-1], s0[prices.size()-1]);
	}
};

int main(void) {
	vector<int> stocks = {1, 2, 3, 0, 2};
	Solution s;
	cout << s.maxProfit(stocks) << endl;
	return 0;
}
