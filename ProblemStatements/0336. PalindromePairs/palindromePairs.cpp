class Solution {
public:
    vector<vector<int>> palindromePairs(vector<string>& words) {
        unordered_map<string, int> id;
        vector<vector<int>> res;
        for (int i {0}; i < words.size(); i++){
            id[words[i]] = i;
        }
        for(int i {0}; i < words.size(); i++){
            string temp = words[i];
            int n = temp.size();
            for(int j {0}; j <= n; j++){
                string pre = temp.substr(0, j);
                string suf = temp.substr(j);
                if(isPalin(pre)){
                    string tar = suf;
                    reverse(tar.begin(), tar.end());
                    auto it = id.find(tar);
                    if(it != id.end() && it -> second != i){
                        res.push_back({it->second, i});
                    }
                }
                if(j != n && isPalin(suf)){
                    string tar = pre;
                    reverse(tar.begin(), tar.end());
                    auto it = id.find(tar);
                    if(it != id.end() && it->second != i){
                        res.push_back({i, it->second});
                    }
                }
            }
        }
        return res;
    }

    bool isPalin(string& s){
        int l {0}, r = s.length() - 1;
        while(l < r){
            if(s[l] != s[r]){
                return false;
            }
            l++;
            r--;
        }
        return true;
    }
};