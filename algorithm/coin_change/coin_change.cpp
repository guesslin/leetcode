#include <iostream>
#include <algorithm>
#include <vector>

using std::vector;
using std::cout;
using std::endl;

class Solution {
	int min(int x, int y) { return x < y? x : y; }
public:
	int coinChange(vector<int>& coins, int amount) {
		int MAX = amount + 1;
		vector<int> count(MAX, MAX);
		count[0] = 0;
		for(int i = 1; i <= amount; ++i) {
			for(auto ite:coins) {
				if(ite <= i) {
					count[i] = min(count[i],  (count[i - ite] + 1));
				}
			}
		}
		return count[amount] > amount ? -1 : count[amount];
	}
};

int main() {
	vector<int> coins;
	coins.push_back(186);
	coins.push_back(419);
	coins.push_back(83);
	coins.push_back(408);
	Solution s;
	cout<<s.coinChange(coins, 6249)<<endl;
	return 0;
}
