class Solution {
public:
    long long minimumCost(string source, string target, vector<char>& original, vector<char>& changed, vector<int>& cost) {
        using ll = long long;
        ll inf = INT_MAX;
        vector<vector<ll>> dist(26, vector<ll>(26,  inf));
        for(int i {0}; i < 26; i++){
            dist[i][i] = 0;
        }
        int m = original.size();
        for(int i {0}; i < m; i++){
            int u = original[i] - 'a';
            int v = changed[i] - 'a';
            dist[u][v] = min(dist[u][v], (ll)cost[i]);
        }
        for(int i {0}; i <26; i++){
            for(int j {0}; j < 26; j++){
                if(dist[j][i] >= inf){
                    continue;
                }
                for(int k {0}; k < 26; k++){
                    dist[j][k] = min(dist[j][k], dist[j][i] + dist[i][k]);
                }
            }
        }
        ll res {0};
        int n = source.length();
        for(int i {0}; i < n;i++){
            if(source[i] == target[i]){
                continue;
            }
            int u = source[i] - 'a';
            int v = target[i] - 'a';
            if(dist[u][v] >= inf){
                return -1;
            }
            res += dist[u][v];
            // cout << dist[u][v] << " " << res << endl;
        }
        return res;
    }
};