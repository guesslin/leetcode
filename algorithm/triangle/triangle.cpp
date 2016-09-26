#include <iostream>
#include <vector>
#include <algorithm>

using std::min;
using std::vector;

class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int levels = triangle.size() - 1;
        for (int row = 1; row <= levels;row++) {
            triangle[row][0] += triangle[row-1][0];
            for (int col = 1; col < triangle[row].size()-1; col++) {
                triangle[row][col] += min(triangle[row-1][col], triangle[row-1][col-1]);
            }
            triangle[row][triangle[row].size()-1] += triangle[row-1][triangle[row-1].size()-1];
        }
        int small = triangle[levels][0];
        for (int col = 1; col < triangle[levels].size(); col++) {
            small = min(small, triangle[levels][col]);
        }
        return small;
    }
};
