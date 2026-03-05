// class Solution {
// public:
//     using ll = long long;
//     using pii = pair<ll, ll>;
//     using T = tuple<ll, ll, int>;

//     ll dp[1 << 12];
//     vector<ll> sizes, medians;
//     vector<vector<int>> lr;

//     ll h1(int mask){
//         ll s {0};
//         for(int i {0}; i < sizes.size(); i++){
//             if(mask & (1 << i)){
//                 s += sizes[i];
//             }
//         }
//         return s;
//     }

//     ll h2(int mask){
//         vector<ll> mer;
//         for(int i{0}; i < sizes.size(); i++){
//             if(mask & (1 << i)){
//                 mer.insert(mer.end(), lr[i].begin(), lr[i].end());
//             }
//         }
//         nth_element(mer.begin(), mer.begin() + mer.size()/2, mer.end());
//         return mer[mer.size() / 2];
//     }

//     long long minMergeCost(vector<vector<int>>& lists) {
//         // priority_queue<T, vector<T>, greater<T>> pq;
//         // for(auto& i : lists){
//         //     ll s = i.size();
//         //     ll med = i[s / 2];
//         //     pq.emplace(s, med, s);
//         // }
//         // ll res {0};
//         // while(pq.size() > 1){
//         //     auto[cost1, med1, s1] = pq.top();
//         //     pq.pop();
//         //     auto[cost2, med2, s2] = pq.top();
//         //     pq.pop();
//         //     ll m = s1 + s2 + abs(med1 - med2);
//         //     res += m;
//         //     ll ns = s1 + s2;
//         //     pq.emplace(ns, (med1 * s1 + med2 * s2) / ns, ns);
//         // }
//         // return res;
//         int n = lists.size();
//         sizes.resize(n);
//         medians.resize(n);
//         lr = lists;
//         for(int i {0}; i < n; i++){
//             sizes[i] = lists[i].size();
//             medians[i] = lists[i][sizes[i] / 2];
//         }
//         // memset(dp, 0x3f, sizeof(dp));
//         // dp[0]= 0;
//         fill(dp, dp + (1 << n), (1LL << 60));
//         dp[0] = 0;
//         for(int i = 0; i < n; i++) {
//             dp[1 << i] = 0;
//         }
//         for(int mask = 1; mask < (1 << n); mask++){
//             int bits = __builtin_popcount(mask);
//             if(bits <= 1){
//                 continue;
//             }
//             for(int sub = mask; sub > 0; sub = (sub - 1) & mask){
//                 if(sub == mask || sub == 0){
//                     continue;
//                 }
//                 int comp = mask ^ sub;
//                 if(__builtin_popcount(sub) == 0 || __builtin_popcount(comp) == 0){
//                     continue;
//                 }
//                 ll s1 = h1(sub);
//                 ll med1 = h2(sub);
//                 ll s2 = h1(comp);
//                 ll med2 = h2(comp);
//                 ll m = s1 + s2 + abs(med1 - med2);
//                 if(dp[sub] < (1LL << 59) && dp[comp] < (1LL << 59)){
//                     dp[mask] = min(dp[mask], dp[sub] + dp[comp] + m);
//                 }
//             }
//         }
//         return dp[(1 << n) - 1];
//     }
// };

class Solution {
public:
    using ll = long long;
    using pii = pair<ll, ll>;
    using T = tuple<ll, ll, int>;

    ll dp[1 << 12];
    vector<ll> length, med;
    vector<vector<int>> lr;

    long long minMergeCost(vector<vector<int>>& lists) {
        int n = lists.size();
        int full_mask = (1 << n) - 1;        
        length.resize(1 << n, 0);
        med.resize(1 << n, 0);
        lr = lists;
        for(int mask = 1; mask < (1 << n); mask++) {
            vector<int> elements;
            for(int i = 0; i < n; i++) {
                if(mask & (1 << i)) {
                    elements.insert(elements.end(), lists[i].begin(), lists[i].end());
                }
            }
            sort(elements.begin(), elements.end());
            length[mask] = elements.size();
            med[mask] = elements[(elements.size() - 1) / 2];
        }        
        fill(dp, dp + (1 << n), 1LL << 60);        
        for(int mask = 1; mask < (1 << n); mask++) {
            if((mask & (mask - 1)) == 0) {
                dp[mask] = 0;
                continue;
            }            
            ll best = 1LL << 60;
            for(int submask = (mask - 1) & mask; submask > 0; submask = (submask - 1) & mask){
                int other = mask ^ submask;
                if(submask < other) {
                    ll cost = dp[submask] + dp[other] + length[submask] + length[other] + abs(med[submask] - med[other]);
                    if(cost < best){
                        best = cost;
                    }
                }
            }            
            dp[mask] = best;
        }        
        return dp[full_mask];
    }
};
