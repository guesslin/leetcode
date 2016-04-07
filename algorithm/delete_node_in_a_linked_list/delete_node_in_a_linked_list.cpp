#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    void deleteNode(ListNode* node) {
        if(node==NULL) {
            return;
        }
        ListNode *temp = node->next;
        if(temp != NULL) {
            node->val = temp->val;
            node->next = temp->next;
            delete temp;
        }
    }
};

int main(int argc, char **argv) {
	return 0;
}
