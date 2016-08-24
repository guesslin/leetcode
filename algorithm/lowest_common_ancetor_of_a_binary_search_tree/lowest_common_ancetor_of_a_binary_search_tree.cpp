#include <iostream>

using std::cout;
using std::endl;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
	TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
		if(p->val > q->val) {
			TreeNode *tmp = p;
			p = q;
			q = tmp;
		}
		while(p->val > root->val || root->val > q->val) {
			if (q->val > root->val) {
				root = root->right;
			} else {
				root = root->left;
			}
		}
		return root;
	}
};

int main(void) {
	int a = 2;
	int b = 3;
	int c = a > b ? a : b;
	cout << c << endl;
	return 0;
}
