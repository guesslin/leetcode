#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    bool hasCycle(ListNode *head) {
        if(head == NULL) {
		    return 0;
	    }
        ListNode *oneStep, *twoStep;
        oneStep = head;
        twoStep = head->next;
        while(1) {
            if(twoStep != NULL && twoStep->next != NULL) {
		    twoStep = twoStep->next->next;
	        } else {
		       return 0;
	        }
	        oneStep = oneStep->next;
	        if(oneStep == twoStep) {
		        return 1;
	        }
	    }
    }
};
