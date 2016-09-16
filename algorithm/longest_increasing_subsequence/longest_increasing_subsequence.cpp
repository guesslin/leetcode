#include <iostream>
#include <vector>
#include <algorithm>

using std::vector;
using std::cout;
using std::endl;
using std::max;

class Solution {
public:
	int lengthOfLIS(vector<int>& nums) {
		vector<int> maxLen(nums.size(), 1);
		int len = nums.size();
		int LIS = 0;
		for (int i = 0; i < len; ++i) {
			for (int j = i+1; j < len; ++j) {
				if (nums[i] < nums[j]) {
					maxLen[j] = max(maxLen[j], maxLen[i] + 1);
				}
			}
		}
		for (int i = 0; i < len; ++i) {
			cout << maxLen[i] << "\t";
			LIS = max(LIS, maxLen[i]);
		}
		cout << endl;
		return LIS;
	}
};

int main(void) {
	vector<int> nums;
	nums.push_back(10);
	nums.push_back(9);
	nums.push_back(2);
	nums.push_back(5);
	nums.push_back(3);
	nums.push_back(7);
	nums.push_back(101);
	nums.push_back(18);
	Solution s;
	cout << s.lengthOfLIS(nums) << endl;
	return 0;
}
