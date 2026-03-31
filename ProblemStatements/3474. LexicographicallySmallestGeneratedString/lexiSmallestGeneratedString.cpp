class Solution {
public:
    string generateString(string str1, string str2) {
        int n = str1.size();
        int m = str2.size();
        int l = n + m - 1;
        string res(l, '?');
        vector<bool> fix(l, false);
        for(int i{0}; i < n; i++){
            if(str1[i] == 'T'){
                for(int j {0}; j < m;j++){
                    int p = i + j;
                    if(res[p] != '?' && res[p] != str2[j]){
                        return "";
                    }
                    res[p] = str2[j];
                    fix[p] = true;
                }
            }
        }
        for(int i{0}; i < n; i++){
            if(str1[i] == 'F'){
                bool fg = true;
                for(int j{0}; j < m; j++){
                    if(res[i + j] == '?' || res[i+j] != str2[j]){
                        fg = false;
                        break;
                    }
                }
                if(fg){
                    return "";
                }
            }
        }
        for(char& ch : res){
            if (ch == '?'){
                ch = 'a';
            }
        }
        for(int i {0}; i < n; i++){
            if(str1[i] != 'F'){
                continue;
            }
            bool fg = true;
            for(int j{0}; j < m; j++){
                if(res[i+j] != str2[j]){
                    fg = false;
                    break;
                }
            }
            if(!fg){
                continue;
            }
            int pos {-1};
            for(int j = m - 1; j >= 0; j--){
                int p = i + j;
                if(!fix[p]){
                    pos = p;
                    break;
                }
            }
            if(pos == -1){
                return "";
            }
            res[pos] = 'b';
            fix[pos] = true;
        }
        return res;
    }
};
