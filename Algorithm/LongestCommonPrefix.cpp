#include <bits/stdc++.h>
using namespace std;

class Solution {
	public:
		bool isCommon(vector<string>& strs, char c, int index) {
			int common = 1;
			for (int i = 1; i < strs.size(); ++i) {
				string str = strs[i];
				if (str.size() <= index) return false;
				if (str[index] == c) {
					common++;
				}
			}
			return common == strs.size();
		}
		string longestCommonPrefix(vector<string>& strs) {
			int ret = 0;
			if (strs.size() == 0) return "";
		    for (int i = 0; i < strs[0].size(); ++i) {
		    	string str = strs[0];
		    	if (isCommon(strs, str[i], i)) ret += 1;
		    	else break;
		    }
		    string res = "";
		    for (int i = 0; i < ret; ++i) {
		    	res += strs[0][i];
		    }
		    return res;
		}
};


int main(int argc, char const *argv[])
{	string strs[] = {"aca","cba"};
	std::vector<string> v(strs, strs + 2);
	Solution sltn;
	string s = sltn.longestCommonPrefix(v);
	cout << s <<endl;
	return 0;
}
