#include <iostream>
#include <vector>
#include <algorithm>

using std::vector;
using std::cout;
using std::endl;
using std::max;

void print(vector<vector<int>>& nums) {
	for (unsigned int row = 0; row < nums.size(); ++row) {
		for (unsigned int col = 0; col < nums[row].size(); ++col) {
			cout << nums[row][col] << "\t";
		}
		cout << endl;
	}
	cout << endl;
}

class Solution {
public:
	int maxCoins(vector<int>& nums) {
		int n = nums.size()+2;
		vector<int> coins (n, 1);
		for (unsigned int i = 0; i < nums.size(); i++) {
			coins[i+1] = nums[i];
		}
		vector<vector<int>> counts (n, vector<int>(n, 0));
		/* k 表示可能的間隔長度 */
		for (int k = 2; k < n; ++k) {
			/* left 表示可以選擇的所有左元素 */
			for (int left = 0; left < n - k; ++left) {
				/* right 表示被選定當前間隔下的右元素 */
				int right = left + k;
				/* i 表示左右元素之間可能被選定的氣球 */
				for (int i = left+1; i < right; ++i) {
					counts[left][right] = max(counts[left][right], coins[left]*coins[i]*coins[right]+counts[left][i]+counts[i][right]);
					cout << left << "\t" << i << "\t" << right << endl;
					print(counts);
				}
			}
		}
		return counts[0][n-1];
	}
};


int main(void) {
	vector<int> coins{3, 1, 5, 8};
	Solution s;
	cout << s.maxCoins(coins) << endl;
	return 0;
}
