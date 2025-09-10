class Solution {
public:
    string removeKdigits(string num, int k) {
        if (k >= num.size()){
            return "0";
        }
        string st;
        st.reserve(num.size());
        for (char c : num){
            while(k > 0 && !st.empty() && st.back() > c){
                st.pop_back();
                --k;
            }
            st.push_back(c);
        }
        while(k > 0 && !st.empty()){
            st.pop_back();
            k--;
        }
        int i {0};
        while(i < st.size() && st[i] == '0'){
            i++;
        }
        string res = st.substr(i);
        return res.empty() ? string("0") : res;
    }
};