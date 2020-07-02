#include <bits/stdc++.h>

using namespace std;

/* Memory Usage: 6.2 MB, less than 100.00% of C++ online submissions for Valid Parentheses. */

bool isValid(string s) {
  stack <char> st;
  for(int i = 0; i < s.size(); i++) {
    char _s = s[i];
    bool flag = false;
    if (_s == ')' && st.size() && st.top() == '(' ) {
      flag = true;
      st.pop();
    }
    if (_s == '}' && st.size() && st.top() == '{' ) {
      flag = true;
      st.pop();
    }
    if (_s == ']' && st.size() && st.top() == '[' ) {
      flag = true;
      st.pop();
    }
    if (!flag) {
      st.push(_s);
    }
  }
  return st.size()==0;
}

int main () {
  cout << isValid(")");
  return 0;
}