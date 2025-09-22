class MovieRentingSystem {
public:
    unordered_map<int, set<pair<int, int>>> un;
    map<pair<int, int>, int>sm;
    set<tuple<int, int, int>> r;

    MovieRentingSystem(int n, vector<vector<int>>& entries) {
        for(const auto& e: entries){
            int s = e[0];
            int m = e[1];
            int p = e[2];
            un[m].insert({p, s});
            sm[{s, m}] = p;
        }
    }
    
    vector<int> search(int movie) {
        vector<int> res;
        auto it = un.find(movie);
        if(it != un.end()){
            int count {0};
            for(auto& [p, s] : it->second){
                if(count >= 5){
                    break;
                }
                res.push_back(s);
                count++;
            }
        }
        return res;
    }
    
    void rent(int shop, int movie) {
        int p = sm[{shop, movie}];
        un[movie].erase({p, shop});
        r.insert({p, shop, movie});
    }
    
    void drop(int shop, int movie) {
        int p = sm[{shop, movie}];
        un[movie].insert({p, shop});
        r.erase({p, shop, movie});
    }
    
    vector<vector<int>> report() {
        vector<vector<int>> res;
        int count {0};
        for(auto& [p, s, m] : r){
            if(count >= 5){
                break;
            }
            res.push_back({s, m});
            count++;
        }
        return res;
    }
};

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * MovieRentingSystem* obj = new MovieRentingSystem(n, entries);
 * vector<int> param_1 = obj->search(movie);
 * obj->rent(shop,movie);
 * obj->drop(shop,movie);
 * vector<vector<int>> param_4 = obj->report();
 */