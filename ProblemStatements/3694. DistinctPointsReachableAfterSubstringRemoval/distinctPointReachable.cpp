class Solution {
public:
    struct p_hash{
        inline size_t operator()(const pair<int, int>& v) const{
            return v.first * 31 + v.second;
        }
    };

    int distinctPoints(string s, int k) {
        int n = s.length();
        int x {0}, y {0};
        for(char c : s){
            if(c == 'U'){
                y++;
            } else if(c == 'D'){
                y--;
            } else if(c == 'L'){
                x--;
            } else if(c == 'R'){
                x++;
            }
        }
        unordered_set<pair<int, int>, p_hash> d;
        int wx {0}, wy {0};
        for(int i {0}; i < k; i++){
            if(s[i] == 'U'){
                wy++;
            } else if(s[i] == 'D'){
                wy--;
            } else if(s[i] == 'L'){
                wx--;
            } else if(s[i] == 'R'){
                wx++;
            }
        }
        d.insert({x - wx, y - wy});
        for(int i {0}; i < n - k; i++){
            char out = s[i];
            char in = s[i + k];
            if(out == 'U'){
                wy--;
            } else if(out == 'D'){
                wy++;
            } else if(out == 'L'){
                wx++;
            } else if(out == 'R'){
                wx--;
            }
            if(in == 'U'){
                wy++;
            } else if(in == 'D'){
                wy--;
            } else if(in == 'L'){
                wx--;
            } else if(in == 'R'){
                wx++;
            }
            d.insert({x - wx, y - wy});
        }
        return d.size();
    }
};