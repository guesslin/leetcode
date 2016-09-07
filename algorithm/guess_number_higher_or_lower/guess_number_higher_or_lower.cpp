#include <iostream>

// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
int guess(int num);

class Solution {
	public:
		int guessNumber(int n) {
			int lo = 1;
			int hi = n;
			int res = 0;
			while(lo < hi) {
				int mid = lo + (hi-lo) / 2;
				res = guess(mid);
				if (res == 0) {
					return mid;
				} else if (res == 1) {
					lo = mid + 1;
				} else {
					hi = mid - 1;
				}
			}
			return lo;
		}
};
