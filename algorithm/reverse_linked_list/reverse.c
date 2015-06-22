#include <stdio.h>

struct ListNode {
	int val;
	struct ListNode *next;
};

struct ListNode* reverse(struct ListNode* head, struct ListNode* pre) {
	if (head->next == NULL) {
		head->next = pre;
		return head;
	}
	struct ListNode* tmp = head->next;
	head->next = pre;
	return reverse(tmp, head);

}

struct ListNode* reverseList(struct ListNode* head) {
	if (head == NULL) {
		return NULL;
	}
	return reverse(head, NULL);
}

int main(int argc, const char *argv[])
{
	printf("hello\n");
	return 0;
}

