#include<iostream>
#include<string>
#include<vector>
#include<stack>
using namespace std;

int scoreOfParentheses(string s){
    int n = s.length();
    stack<int> st;
    //(()(()))(())
    //(()
    //01
    for (size_t i = 0; i < n; i++)
    {
        if (s[i] == '(')
        {
            st.push(0);
        }
        else{
            int v = st.top();
            if (v == 0)
            {
                st.pop();
                st.push(1);
            }
            else
            {
                int temp = 0;
                while (st.top() != 0)
                {
                    int v = st.top();
                    st.pop();
                    temp = temp + v;
                }
                st.pop();
                temp = temp << 1;
                st.push(temp);
            } 
        }
    }
    int res = 0;
    while (!st.empty())
    {
        int v = st.top();
        st.pop();
        res = res + v;
    }
    return res;
}

int main(int argc, char* argv[]){
    // (())
    //int res = Score("()()");
    //int res = Score("()()");
    //int res = Score("(())");
    int res = scoreOfParentheses("(()(()))(())");//output 7
    cout<<res<<endl;
}