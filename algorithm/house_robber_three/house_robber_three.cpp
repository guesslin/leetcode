#include <iostream>
#include <vector>
#include <algorithm>

using std::max;
using std::vector;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
class Solution {
    vector<int> robSub(TreeNode *root) {
        vector<int> res (2, 0);
        if (root == NULL) {
            return res;
        }
        vector<int> left = robSub(root->left);
        vector<int> right = robSub(root->right);
        res[0] = root->val + left[1] + right[1];
        res[1] = max(max(left[0]+right[0], left[1]+right[1]), max(left[0]+right[1], left[1]+right[0]));
        return res;
        
    }
public:
    int rob(TreeNode* root) {
        vector<int> res = robSub(root);
        return res[0] > res[1] ? res[0] : res[1];
    }
};

// (withRoot, withOutRoot)
// withRoot = root + leftWithOutRoot + rightWithOutRoot
// withOutRoot = max(leftWithRoot+rightWithRoot, leftWithOutRoot+rightWithOutRoot, leftWithRoot+rightWithOutRoot, leftWithOutRoot+rightWithRoot)
