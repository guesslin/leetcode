#include <iostream>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    bool hasPathSum(TreeNode* root, int sum) {
        if (root == NULL) {
		return false;
	}
	if (sum == root->val && (root->left == NULL && root->right == NULL)) {
		return true;
	}
	return hasPathSum(root->right, sum-root->val) || hasPathSum(root->left, sum-root->val);
    }
};
