class Solution {
public:
    struct t{
        string name, city, raw;
        int time, amt, id;
    };

    vector<string> invalidTransactions(vector<string>& transactions) {
        int n = transactions.size();
        vector<t> trans;
        for(int i {0}; i < n; i++)    {
            stringstream ss(transactions[i]);
            string name, ts, as, city;
            getline(ss, name, ',');
            getline(ss, ts, ',');
            getline(ss, as, ',');
            getline(ss, city, ',');
            t temp{name, city, transactions[i], stoi(ts), stoi(as), i};
            trans.push_back(temp);
        }
        vector<bool> iv(n, false);
        for(int i{0}; i < n; i++){
            if(trans[i].amt > 1000){
                iv[i] = true;
            }
        }
        unordered_map<string, vector<int>> nid;
        for(int i{0}; i < n;i++){
            nid[trans[i].name].push_back(i);
        }
        for(auto& p : nid){
            vector<int>& id = p.second;
            for(int i {0}; i < id.size(); i++){
                for(int j {0}; j < id.size(); j++){
                    if(i == j){
                        continue;
                    }
                    auto& t1 = trans[id[i]], &t2 = trans[id[j]];
                    if(t1.city != t2.city && abs(t1.time - t2.time) <= 60){
                        iv[t1.id] = true;
                    }
                }
            }
        }
        vector<string>res;
        for(int i{0}; i < n; i++){
            if(iv[i]){
                res.push_back(trans[i].raw);
            }
        }
        return res;
    }
};