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
        static vector<int> count({0});
        while(count.size() <= unsigned(n)) {
            int m = count.size();
            int currentMin = m;
            for (int i = 1; i*i <= m; i++)
            {
                currentMin = min(currentMin, count[m - i*i] + 1);
            }
            count.push_back(currentMin);
            
        }
        return count[n];
    }
};
