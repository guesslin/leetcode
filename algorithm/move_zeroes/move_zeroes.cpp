#include <iostream>
#include <vector>
#include <queue>

using std::vector;
using std::queue;
using std::cout;
using std::endl;

class Solution {
public:
	void moveZeroes(vector<int>& nums) {
		queue<int> nptr;
		for(unsigned int i = 0; i < nums.size(); ++i) {
			if(nums[i] == 0) {
				nptr.push(i);
			} else {
				if(!nptr.empty()) {
					nums[nptr.front()] = nums[i];
					nptr.pop();
					nptr.push(i);
					nums[i] = 0;
				}
			}
		}
	}
};

int main(int argc, char **argv) {
	Solution s;
	vector<int> nums (10, 0);
	nums[0] = 4;
	nums[1] = 2;
	nums[2] = 4;
	nums[5] = 3;
	nums[7] = 5;
	nums[8] = 1;
	s.moveZeroes(nums);
	for(auto ite:nums) {
		cout<<ite<<" ";
	}
	cout<<endl;
	return 0;
}
