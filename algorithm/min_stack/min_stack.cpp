#include <iostream>
#include <stack>

using std::stack;

class MinStack {
private:
    stack<int> data;
    stack<int> minData;
public:
    /** initialize your data structure here. */
    MinStack() {
    }
    
    void push(int x) {
        data.push(x);
        if(minData.empty()){
            minData.push(x);
        } else {
            minData.push(x < minData.top() ? x : minData.top());
        }
    }
    
    void pop() {
        if(!data.empty()) {
            data.pop();
            minData.pop();
        }
    }
    
    int top() {
        return data.top();
    }
    
    int getMin() {
        return minData.top();
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */
