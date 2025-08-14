class Solution {
public:
    vector<vector<string>> accountsMerge(vector<vector<string>>& accounts) {
        unordered_map<string, string> mail;
        unordered_map<string, string> parent;
        for(auto& acc : accounts){
            string name = acc[0];
            for(int i {1}; i < acc.size(); i++){
                string email = acc[i];
                mail[email] = name;
                if(parent.find(email) == parent.end()){
                    parent[email] = email;
                }
            }
        }
        for(auto& acc: accounts){
            string fe = acc[1];
            for(int i {2} ; i < acc.size(); i++){
                h1(fe, acc[i], parent);
            }
        }
        unordered_map<string, vector<string>> grps;
        for(auto& p : mail){
            string email = p.first;
            string root = find(email, parent);
            grps[root].push_back(email);
        }
        vector<vector<string>> res;
        for(auto& g : grps){
            vector<string> temp = g.second;
            sort(temp.begin(), temp.end());
            vector<string> acc;
            acc.push_back(mail[temp[0]]);
            for(string s : temp){
                acc.push_back(s);
            }
            res.push_back(acc);
        }
        return res;
    }

    string find(string x, unordered_map<string, string>& parent){
        if(parent[x] != x){
            parent[x] = find(parent[x], parent);
        }
        return parent[x];
    }

    void h1(string x, string y, unordered_map<string, string>& parent){
        string rx = find(x, parent);
        string ry = find(y, parent);
        if(rx != ry){
            parent[rx] = ry;
        }
    }
};