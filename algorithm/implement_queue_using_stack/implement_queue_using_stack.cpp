#include <iostream>
#include <stack>

using std::stack;

class Queue {
private:
    stack<int> a,b;
    // a 正向 => 1, 2, 3, 4, 5
    // b 逆向 => 5, 4, 3, 2, 1
public:
    // Push element x to the back of queue.
    void push(int x) {
	    // 總是 push 到正向的 a 裡
	    a.push(x);
    }

    // Removes the element from in front of queue.
    void pop(void) {
	    int x;
	    // 把所有數字從 a 之中逆向倒進 b 裡面
	    while(!a.empty()) {
		    x = a.top();
		    a.pop();
		    b.push(x);
	    }
	    // 此時 b 的最後一個元素就是 queue 第一個元素
	    b.pop();
	    // 再把 b 倒回 a 裡面
	    while(!b.empty()) {
		    x = b.top();
		    b.pop();
		    a.push(x);
	    }
    }

    // Get the front element.
    int peek(void) {
	    int res, x;
	    // 想辦法取出 a 之中第一個元素
	    while(!a.empty()) {
		    x = a.top();
		    a.pop();
		    b.push(x);
	    }
	    res = b.top();
	    while(!b.empty()) {
		    x = b.top();
		    b.pop();
		    a.push(x);
	    }
	    return res;
    }

    // Return whether the queue is empty.
    bool empty(void) {
        if(a.empty() && b.empty()) {
            return true;
        }
        return false;
    }
};
