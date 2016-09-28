#include <iostream>
#include <queue>

using std::queue;

class Stack {
private:
    queue<int> store;
public:
    // Push element x onto stack.
    void push(int x) {
        store.push(x);
    }

    // Removes the element on top of the stack.
    void pop() {
        int len = store.size();
        for (int i = 0; i < len-1; ++i) {
            int tmp = store.front();
            store.pop();
            store.push(tmp);
        }
        store.pop();
    }

    // Get the top element.
    int top() {
        return store.back();
    }

    // Return whether the stack is empty.
    bool empty() {
        return store.empty();
    }
};
