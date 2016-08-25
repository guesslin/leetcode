#include <iostream>

using std::cout;
using std::endl;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
	    if(head == NULL) {
		    return 0;
	    }
	    ListNode *oneStep, *twoStep;
	    oneStep = head;
	    twoStep = head->next;
	    while(oneStep != twoStep) {
		    if(twoStep != NULL && twoStep->next != NULL) {
			    twoStep = twoStep->next->next;
		    } else {
			    return NULL;
		    }
		    oneStep = oneStep->next;
	    }
	    oneStep = head;
	    twoStep = twoStep->next;
	    while(oneStep != twoStep) {
		    oneStep = oneStep->next;
		    twoStep = twoStep->next;
	    }
	    return oneStep;

    }
};

int main(void) {
	ListNode head(1);
	head.next = new ListNode(2);
	head.next->next = &head;
	Solution s;
	cout << s.detectCycle(&head) << endl;
	return 0;
}
