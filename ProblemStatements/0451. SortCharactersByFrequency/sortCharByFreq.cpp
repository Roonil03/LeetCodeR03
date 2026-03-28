class Solution {
public:
    string frequencySort(string s) {
        unordered_map<char, int> freq;
        for(char c : s){
            freq[c]++;
        }
        int n = s.size();
        vector<vector<char>> buk(n+1);
        for(auto&[ch, c] : freq){
            buk[c].push_back(ch);
        }
        string res;
        res.reserve(n);
        for(int i = n; i >= 1; i--){
            for(char ch : buk[i]){
                res.append(i, ch);
            }
        }
        return res;
    }
};