#include <iostream>
#include <vector>

using std::vector;
using std::cout;
using std::endl;

class Solution {
	int count(vector<long>& sums, int start, int end, int lower, int upper) {
		if (end-start <= 1) {
			return 0;
		}
		int mid = (start + end ) / 2;
		int c = count(sums, start, mid, lower, upper) + count(sums, mid, end, lower, upper);
		int m = mid, n = mid, k = mid, len = 0;
		vector<long> cache(end-start, 0);
		for (int i = start, s = 0; i < mid; i++, s++) {
			while(m<end && sums[m]-sums[i]<lower) m++;
			while(n<end && sums[n]-sums[i]<=upper) n++;
			c += n-m;
			while(k<end && sums[k]<sums[i]) cache[s++] = sums[k++];
			cache[s] = sums[i];
			len = s;
		}
		for (int i = 0; i <= len; i++) {
			sums[start+i] = cache[i];
		}
		return c;
	}
public:
	int countRangeSum(vector<int>& nums, int lower, int upper) {
		int len = nums.size();
		if (len == 0) {
			return 0;
		}
		vector<long> sums(len+1, 0);
		for(int i = 0; i < len; ++i) {
			sums[i+1] = sums[i] + nums[i];
		}
		return count(sums, 0, len+1, lower, upper);
	}
};

int main(int argc, char **argv) {
	vector<int> nums;
	Solution s;
	nums.push_back(-2);
	nums.push_back(5);
	nums.push_back(-1);
	nums.push_back(1);
	cout<<s.countRangeSum(nums, -2, 2)<<endl;
	return 0;
}
