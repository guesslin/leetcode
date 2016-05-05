#include <iostream>
#include <vector>

using std::vector;
using std::cout;
using std::endl;

class NumArray {
	vector<int> rangeSum;
public:
	NumArray(vector<int> &nums) {
		rangeSum.push_back(0);
		for(auto ite:nums) {
			rangeSum.push_back(rangeSum.back()+ite);
		}

	}

	int sumRange(int i, int j) {
		return rangeSum[j+1] - rangeSum[i];

	}
};

int main() {
	vector<int> nums={1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
	NumArray numArray(nums);
	cout<<numArray.sumRange(0, 1)<<endl;
	cout<<numArray.sumRange(1, 2)<<endl;
	return 0;
}


// Your NumArray object will be instantiated and called as such:
// NumArray numArray(nums);
// numArray.sumRange(0, 1);
// numArray.sumRange(1, 2);
