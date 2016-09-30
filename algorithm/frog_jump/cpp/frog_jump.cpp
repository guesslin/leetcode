#include <iostream>
#include <map>
#include <vector>
#include <set>

using namespace std;

class Solution {
public:
    bool canCross(vector<int>& stones) {
	    if (stones.size() <= 1) {
		    return true;
	    }
	    map<int, map<int, bool>> stonesMap;
	    for (auto ite:stones) {
		    stonesMap[ite] = map<int, bool>();
	    }
	    stonesMap[1][1] = true;
	    for (auto s : stones) {
		    for (map<int,bool>::iterator it=stonesMap[s].begin(); it!=stonesMap[s].end(); ++it) {
			    int k = it->first;
			    try {
				    if (k - 1 > 0) {
					    stonesMap[s+k-1][k-1] = true;
				    }
			    } catch (...) {}
			    try {
				    stonesMap[s+k][k] = true;
			    } catch (...) {}
			    try {
				    stonesMap[s+k+1][k+1] = true;
			    } catch (...) {}
		    }
	    }
	    return stonesMap[*(stones.end()-1)].size() > 0;
    }
};

int main(void) {
	vector<int> example1 = {0,1,3,5,6,8,12,17},
		    example2 = {0,1,2,3,4,8,9,11};
	Solution s;
	cout << (s.canCross(example1) ? "True" : "False" ) << endl;
	cout << (s.canCross(example2) ? "True" : "False" ) << endl;
	return 0;
}
