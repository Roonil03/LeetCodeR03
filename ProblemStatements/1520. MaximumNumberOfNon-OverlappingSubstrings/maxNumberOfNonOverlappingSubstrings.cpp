class Solution {
public:
    vector<string> maxNumOfSubstrings(string s) {
        int n = s.size();
        vector<int> l(26, n), r(26, -1);
        for(int i {0}; i < n; i++){
            l[s[i] - 'a'] = min(l[s[i] - 'a'], i);
            r[s[i] - 'a'] = max(r[s[i] - 'a'], i);
        }
        vector<pair<int, int>> v;
        for(int i {0}; i < 26; i++){
            if(l[i] == n){
                continue;
            }
            int b = l[i], e = r[i];
            bool fg = true;
            for(int j = b;  j <= e; j++){
                if(l[s[j] - 'a'] < b){
                    fg = false;
                    break;
                }
                e = max(e, r[s[j] - 'a']);
            }
            if(fg){
                v.emplace_back(e, b);
            }
        }
        ranges::sort(v, [](auto& a, auto& b){
            return a.first == b.first ? a.second > b.second : a.first < b.first;
        });
        vector<string> res;
        int ll = -1;
        for(auto& [e, b] : v){
            if(b > ll){
                res.push_back(s.substr(b, e - b + 1));
                ll = e;
            }
        }
        return res;
    }
};