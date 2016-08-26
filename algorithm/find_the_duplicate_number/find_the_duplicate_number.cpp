#include <iostream>
#include <vector>

/*
 * You must not modify the array (assume the array is read only).
 * You must use only constant, O(1) extra space.
 * Your runtime complexity should be less than O(n^2).
 * There is only one duplicate number in the array, but it could be repeated more than once.
 */

using std::vector;
using std::cout;
using std::endl;

class Solution {
public:
	int findDuplicate(vector<int>& nums) {
		int low = 0;
		int up = nums.size() - 1;
		int mid = (low+up)/2;
		int count;
		// use Pigeonhole priciple
		while(low < up) {
			count = 0;
			for(auto i : nums) {
				if (i <= mid) {
					++count;
				}
			}
			if (count > mid) {
				up = mid;
			} else {
				low = mid+1;
			}
			if (mid == (up+low)/2) {
				mid = (up+low)/2+1;
			} else {
				mid = (up+low)/2;
			}
		}
		return low;
	}
};

int main(void) {
	vector<int> nums {2, 2, 2, 2, 2};
	Solution s;
	cout << s.findDuplicate(nums) << endl;
}
