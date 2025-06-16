class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        vector<int>res;
        if(s.size() < p.size()){
            return res;
        }
        vector<int>tar(26,0), win(26, 0);
        for(char c : p){
            tar[c-'a']++;
        }
        int pl = p.size();
        for(int i {0}; i < pl; i++){
            win[s[i] - 'a']++;
        }
        if(win == tar){
            res.push_back(0);
        }
        for(int i = pl; i < s.size(); i++){
            win[s[i]-'a']++;
            win[s[i-pl] - 'a']--;
            if(win == tar){
                res.push_back(i - pl + 1);
            }
        }
        return res;
    }
};