class Solution {
    using ll = long long;
public:
    vector<long long> distance(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> mp;
        for (int i {0}; i < n; i++){
            mp[nums[i]].push_back(i);
        }
        vector<ll> res(n, 0);
        for(auto& i : mp){
            vector<int>& id = i.second;
            int k = id.size();
            vector<ll> pre(k);
            pre[0] = id[0];
            for(int i {1}; i < k; i++){
                pre[i] = pre[i-1] + id[i];
            }
            for(int j {0}; j < k; j++){
                ll l {0}, r{0};
                if(j > 0){
                    l = (ll)j * id[j] - pre[j-1];
                }
                if(j < k - 1){
                    r = (pre[k-1] - pre[j]) - (ll)(k - j - 1) * id[j];
                }
                res[id[j]] = l + r;
            }
        }
        return res;
    }
};