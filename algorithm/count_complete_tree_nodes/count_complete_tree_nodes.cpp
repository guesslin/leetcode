#include <cmath>
#include <iostream>

using std::pow;
using std::cout;
using std::endl;

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

bool isExist(TreeNode *root, int depth) {
	for (; depth > 0; --depth) {
		if (root == NULL) {
			return false;
		}
		root = root->left;
	}
	return true;
}

class Solution {
public:
	int countNodes(TreeNode *root) {
		if (root == NULL) {
			return 0;
		}
		int depth = 1;
		auto cur = root;
		while (cur->left != NULL) {
			++depth;
			cur = cur->left;
		}
		int lo = pow(2, depth-1), last = 0, cdepth = 0;
		for (cur = root; cdepth < depth-1; ++ cdepth) {
			if (isExist(cur->right, depth - cdepth - 1)) {
				cur = cur->right;
				last = last*2+1;
			} else {
				cur = cur->left;
				last = last*2;
			}
		}
		return lo+last;
	}
};

int main(void) {
	TreeNode one {1};
	TreeNode two {2};
	TreeNode three {3};
	TreeNode four {4};
	TreeNode five {5};
	TreeNode six {6};
	TreeNode seven {7};
	TreeNode eight {8};
	TreeNode nine {9};
	TreeNode ten {10};
	TreeNode eleven {11};
	TreeNode twelve {12};
	one.left = &two;
	one.right = &three;
	two.left = &four;
	two.right = &five;
	three.left = &six;
	three.right = &seven;
	four.left = &eight;
	four.right = &nine;
	five.left = &ten;
	five.right = &eleven;
	six.left = &twelve;
	Solution s;
	cout << s.countNodes(&one) << endl;
	return 0;
}
