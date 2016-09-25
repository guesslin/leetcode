#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* oddEvenList(ListNode* head) {
        if (head == NULL) {
            return NULL;
        }
        auto odd = head;
        auto even = head->next;
        auto evenHead = head->next;
        while(even != NULL && odd != NULL) {
            odd->next = even->next;
            if (odd->next == NULL) {
                break;
            }
            odd = odd->next;
            even->next = odd->next;
            if (even->next == NULL) {
                break;
            }
            even = even->next;
        }
        odd -> next = evenHead;
        return head;
    }
};
