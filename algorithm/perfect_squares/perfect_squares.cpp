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
        int maxSquare = int(sqrt(n));
        vector<int> count(n+1, n+1);
        vector<int> squares(maxSquare);
        for(int i = 0; i < maxSquare; i++) {
            squares[i] = (i+1)*(i+1);
        }
        count[0] = 0;
        for(int i = 1; i <= n; i++){
            for(auto ite:squares){
                if(ite <= i) {
                    count[i] = min(count[i], count[i-ite]+1);
                }
            }
        }
        return count[n];
    }
};
