#include <iostream>
#include <vector>

using std::vector;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
 
class Solution {
public:
    bool isBalanced(TreeNode* root) {
        if (root == NULL) {
            return true;
        }
        int ld = depth(root->left);
        int rd = depth(root->right);
        if (ld == -1 || rd == -1 || abs(rd-ld) > 1) {
            return false;
        }
        return true;
        
    }
    
    int depth(TreeNode* node) {
        if (node == NULL) {
            return 0;
        }
        int ld = depth(node->left);
        int rd = depth(node->right);
        if (ld == -1 || rd == -1 || abs(rd-ld) > 1) {
            return -1;
        }
        return ld > rd ? ld+1 : rd+1;
    }
    
    int abs(int x) {
        if (x < 0) {
            return -x;
        }
        return x;
    }
};
