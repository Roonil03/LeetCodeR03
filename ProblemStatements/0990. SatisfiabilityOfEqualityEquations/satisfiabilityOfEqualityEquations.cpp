class Solution {
public:
    bool equationsPossible(vector<string>& equations) {
        vector<int> par(26);
        for (int i {0}; i < 26; i++){
            par[i] = i;
        }
        for(string& eq : equations){
            if(eq[1] == '='){
                int x = eq[0] - 'a';
                int y = eq[3] - 'a';
                h1(x, y, par);
            }
        }
        for(string& eq : equations){
            if(eq[1] == '!'){
                int x = eq[0] - 'a';
                int y = eq[3] - 'a';
                if(find(x, par) == find(y, par)){
                    return false;
                }
            }
        }
        return true;
    }

    int find(int x, vector<int>& par){
        if(par[x] != x){
            par[x] = find(par[x], par);
        }
        return par[x];
    }

    void h1(int x, int y, vector<int>& par){
        int rx = find(x, par);
        int ry = find(y, par);
        if(rx != ry){
            par[rx] = ry;
        }
    }
};