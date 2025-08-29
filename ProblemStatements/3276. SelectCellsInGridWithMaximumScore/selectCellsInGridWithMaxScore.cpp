// class Solution {
// public:
//     int maxScore(vector<vector<int>>& grid) {
//         // int n = grid.size(), m = grid[0].size();
//         // int res {0};
//         // unordered_set<int>u;
//         // function<void(int, int)> bt = [&](int r, int cur){
//         //     if(r == n){
//         //         res = max(res, cur);
//         //         return;
//         //     }
//         //     bt(r + 1 , cur);
//         //     for(int c {0}; c < m; c++){
//         //         if(u.find(grid[r][c]) == u.end()){
//         //             u.insert(grid[r][c]);
//         //             bt(r + 1, cur + grid[r][c]);
//         //             u.erase(grid[r][c]);
//         //         }
//         //     }
//         // };
//         // bt(0, 0);
//         // return res;

//         int n = grid.size(), m = grid[0].size();
//         unordered_map<int,int> mp;
//         int id {0};
//         for(int i {0}; i < n; i++){
//             for(int j {0}; j < m; j++){
//                 if(mp.find(grid[i][j]) == mp.end()){
//                     mp[grid[i][j]] = id++;
//                 }
//             }
//         }
//         int t = id;
//         if(t > 20){
//             return bt(grid);
//         }
//         vector<int> dp(1 << t, -1);
//         dp[0] = 0;
//         for(int i {0}; i < n; i++){
//             vector<int> nd = dp;
//             for(int mk {0}; mk < (1<<t); mk++){
//                 if(dp[mk] == -1){
//                     continue;
//                 }
//                 for(int j {0}; j < m; j++){
//                     int v = grid[i][j];
//                     int ix = mp[v];
//                     if((mk & (1 << ix)) == 0){
//                         int nm = mk | (1<<ix);
//                         nd[nm] = max(nd[nm], dp[mk] + v);
//                     }
//                 }
//             }
//             dp = nd;
//         }
//         return *max_element(dp.begin(), dp.end());
//     }

//     int bt(vector<vector<int>>& g) {
//         int n = g.size(), m = g[0].size();
//         int rs {0};
//         unordered_set<int> us;        
//         function<void(int, int)> h1 = [&](int r, int c) {
//             if(r == n) {
//                 rs = max(rs, c);
//                 return;
//             }
//             if(c + (n - r) * 100 <= rs){
//                 return;
//             }
//             h1(r + 1, c);
//             for(int j {0}; j < m; j++){
//                 if(us.count(g[r][j]) == 0){
//                     us.insert(g[r][j]);
//                     h1(r + 1, c + g[r][j]);
//                     us.erase(g[r][j]);
//                 }
//             }
//         };        
//         h1(0, 0);
//         return rs;
//     }
// };

class Solution {
public:
    int dp[2001][101];
    int solve(int i,int n,vector<int>&t,map<int,vector<int>>&mp,int mask) {
        if(i==n)
            return 0;
        if(dp[mask][i]!=-1)
            return dp[mask][i];
        int ans=0;
        for(auto j:mp[t[i]]) {
            if((mask&(1<<j))==0) {
                ans=max(ans,t[i]+solve(i+1,n,t,mp,(mask|(1<<j))));
            }
        }
        ans=max(ans,solve(i+1,n,t,mp,mask));
        return dp[mask][i]=ans;
    }
    int maxScore(vector<vector<int>>& v) {
        int n = v.size();
        int m = v[0].size();
        set<int>s;
        for(int i=0;i<n;i++) {
            for(int j=0;j<m;j++) {
                s.insert(v[i][j]);
            }
        }
        vector<int>t;
        for(auto i:s)
            t.push_back(i);
        sort(t.begin(),t.begin(),greater<int>());
        map<int,vector<int>>mp;
        for(int i=0;i<n;i++) {
            for(int j=0;j<m;j++) {
                mp[v[i][j]].push_back(i);
            }
        }
        memset(dp,-1,sizeof(dp));
        return solve(0,t.size(),t,mp,0);
    }
};