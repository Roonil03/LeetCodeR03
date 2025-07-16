class Solution {
public:
    string removeDuplicateLetters(string s) {
        vector<int>count (26, 0);
        for(char ch : s){
            count[ch - 'a']++;
        }
        vector<bool>a (26, false);
        string res = "";
        for(char ch : s){
            count[ch - 'a']--;
            if(a[ch - 'a']){
                continue;
            }
            while(!res.empty() && res.back() > ch && count[res.back() - 'a'] > 0){
                a[res.back() - 'a'] = false;
                res.pop_back();
            }
            res.push_back(ch);
            a[ch - 'a'] = true;
        }
        return res;
    }
};