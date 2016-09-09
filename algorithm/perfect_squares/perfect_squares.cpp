#include <iostream>
#include <vector>
#include <cmath>
#include <algorithm>

using std::vector;
using std::sqrt;
using std::min;

class Solution {
public:
	int numSquares(int n) {
		vector<int> count(n+1, n+1);
		count[0] = 0;
		count[1] = 1;
		for(int i = 2; i <= n; i++){
			int a = 1;
			while(a*a <= i) {
				count[i] = min(count[i], count[i-a*a]+1);
				a++;
			}
		}
		return count[n];
	}
};
