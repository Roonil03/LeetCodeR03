using ll = long long;

class Solution {
public:
    long long totalCost(vector<int>& costs, int k, int candidates) {
        // int n = costs.size();
        // ll res {0};
        // int l {0}, r = n - 1;
        // priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pqL, pqR;
        // for(int i{0}; i < candidates && l <= r; i++){
        //     pqL.emplace(costs[r], r);
        //     r--;
        // }
        // for(int i{0}; i < k; i++){
        //     if(!pqL.empty() && !pqR.empty()){
        //         if(pqL.top().first < pqR.top().first || (pqL.top().first == pqR.top().first && pqL.top().second < pqR.top().second)){
        //             res += pqL.top().first;
        //             pqL.pop();
        //             if (l <= r){
        //                 pqL.emplace(costs[l], l);
        //                 l++;
        //             }
        //         } else{
        //             res += pqR.top().first;
        //             pqR.pop();
        //             if(l <= r){
        //                 pqR.emplace(costs[r], r);
        //                 r--;
        //             }
        //         }
        //     } else if(!pqL.empty()){
        //         res += pqL.top().first;
        //         pqL.pop();
        //         if (l <= r){
        //             pqL.emplace(costs[l], l);
        //             l++;
        //         }
        //     } else{
        //         res += pqR.top().first;
        //         pqR.pop();
        //         if (l <= r){
        //             pqR.emplace(costs[r], r);
        //             r--;
        //         }
        //     }
        // }
        // return res;
        int i {0};
        int j = costs.size() - 1;
        priority_queue<int, vector<int>, greater<int>> pq1;
        priority_queue<int, vector<int>, greater<int>> pq2;
        ll res {0};
        while(k--){
            while(pq1.size() < candidates && i <= j){
                pq1.push(costs[i++]);
            }
            while(pq2.size() < candidates && i <= j){
                pq2.push(costs[j--]);
            }
            int t1 = pq1.size() > 0 ? pq1.top() : INT_MAX;
            int t2 = pq2.size() > 0 ? pq2.top() : INT_MAX;
            if(t1 <= t2){
                res += t1;
                pq1.pop();
            }
            else{
                res += t2;
                pq2.pop();
            }
        }
        return res;
    }
};