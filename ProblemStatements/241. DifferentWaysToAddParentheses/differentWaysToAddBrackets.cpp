#include <iostream>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;
class Solution{
    public:
vector<int> diffWaysToCompute(string exp) {
    vector<int> v1;

    if (exp.find_first_not_of("0123456789") == string::npos) {
        v1.push_back(stoi(exp));
        return v1;
    }

    for(int i{0}; i<exp.size(); i++) {
        if(exp[i] == '+' || exp[i] == '-' || exp[i] == '*') {
            vector<int> left = diffWaysToCompute(exp.substr(0, i));
            vector<int> right = diffWaysToCompute(exp.substr(i+1));

            for(int l: left) {
                for(int r: right) {
                    if(exp[i] == '+') v1.push_back(l+r);
                    if(exp[i] == '-') v1.push_back(l-r);
                    if(exp[i] == '*') v1.push_back(l*r);
                }
            }
        }
    }
    sort(v1.begin(), v1.end());
    return v1;
}
};

int main() {
    string testing_v1 = "2*3-4*5";
    vector<int> I_AM_SO_ANNOYED_NOW_PLS_HELP_ME = sol(testing_v1);
    for(int j: I_AM_SO_ANNOYED_NOW_PLS_HELP_ME)
        cout << j << ", ";
}