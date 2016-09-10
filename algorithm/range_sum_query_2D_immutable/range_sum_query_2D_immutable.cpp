#include <iostream>
#include <vector>

using std::vector;

class NumMatrix {
    vector<vector<int> > counts;
public:
    NumMatrix(vector<vector<int> > &matrix) {
        int rows = matrix.size();
        if(rows == 0) {
            return;
        }
        int cols = matrix[0].size();
        vector<int> zeroRow(cols+1, 0);
        counts.push_back(zeroRow);
        for(int row = 0; row < rows; row++) {
            vector<int> singleRow(cols+1, 0);
            for(int col = 1; col <= cols; col++) {
                singleRow[col] = singleRow[col-1] + matrix[row][col-1];
            }
            counts.push_back(singleRow);
        }
	for(int col = 1; col <= cols; col++) {
		for(int row = 1; row <= rows; row++) {
			counts[row][col] += counts[row-1][col];
		}
	}
    }

    int sumRegion(int row1, int col1, int row2, int col2) {
        if(counts.size() == 0) {
            return 0;
        }
        return counts[row2+1][col2+1] - counts[row2+1][col1] - counts[row1][col2+1] + counts[row1][col1];
    }
};


int main(void) {
	vector<vector<int> > matrix;
	int arr1[] = {3, 0, 1, 4, 2};
	int arr2[] = {5, 6, 3, 2, 1};
	int arr3[] = {1, 2, 0, 1, 5};
	int arr4[] = {4, 1, 0, 1, 7};
	int arr5[] = {1, 0, 3, 0, 5};
	vector<int> one (arr1, arr1+sizeof(arr1)/sizeof(arr1[0]));
	vector<int> two (arr2, arr2+sizeof(arr2)/sizeof(arr2[0]));
	vector<int> three (arr3, arr3+sizeof(arr3)/sizeof(arr3[0]));
	vector<int> four (arr4, arr4+sizeof(arr4)/sizeof(arr4[0]));
	vector<int> five (arr5, arr5+sizeof(arr5)/sizeof(arr5[0]));
	matrix.push_back(one);
	matrix.push_back(two);
	matrix.push_back(three);
	matrix.push_back(four);
	matrix.push_back(five);
	NumMatrix numMatrix(matrix);
	std::cout << numMatrix.sumRegion(2, 1, 4, 3) << std::endl;
	std::cout << numMatrix.sumRegion(1, 1, 2, 2) << std::endl;
	std::cout << numMatrix.sumRegion(1, 2, 2, 4) << std::endl;
	return 0;
}

// [[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]],sumRegion(2,1,4,3),sumRegion(1,1,2,2),sumRegion(1,2,2,4)

// Your NumMatrix object will be instantiated and called as such:
// NumMatrix numMatrix(matrix);
// numMatrix.sumRegion(0, 1, 2, 3);
// numMatrix.sumRegion(1, 2, 3, 4);
