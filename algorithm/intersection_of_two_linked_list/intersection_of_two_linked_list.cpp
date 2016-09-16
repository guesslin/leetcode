#include <iostream>

struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
		if (headA == NULL || headB == NULL) {
			return NULL;
		}
		ListNode *curA = headA;
		ListNode *curB = headB;
		while(curA != curB) {
			curA = curA == NULL ? headB : curA->next;
			curB = curB == NULL ? headA : curB->next;
		}
		return curA;
	}
};
