class Solution {
public:
    long long maxScore(vector<int>& nums1, vector<int>& nums2, int k) {
        int n = nums1.size();
        vector<pair<int, int>> v(n);
        for (int i {0}; i < n; i++){
            v[i] = {nums2[i], nums1[i]};
        }
        sort(v.rbegin(), v.rend());
        priority_queue<int, vector<int>, greater<int>> pq;
        long long s{0}, res{0};
        for(int i {0}; i < n; i++){
            s += v[i].second;
            pq.push(v[i].second);
            if(pq.size() > k){
                s -= pq.top();
                pq.pop();
            }
            if (pq.size() == k){
                res = max(res, s * v[i].first);
            }
        }
        return res;
    }
};