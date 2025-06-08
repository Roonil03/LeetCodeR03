using ll = long long;

class Solution {
public:
    vector<long long> mostFrequentIDs(vector<int>& nums, vector<int>& freq) {
          int n = nums.size();
          unordered_map<ll, ll> count;
          map<ll, ll>f;
          vector<ll>res(n);
          for(int  i{0}; i < n; i++){
            int id = nums[i], g = freq[i];
            int prev = count[id];
            if (prev > 0){
                f[prev]--;
                if(f[prev] == 0){
                    f.erase(prev);
                }
            }
            count[id] += g;
            if(count[id] > 0){
                f[count[id]]++;
            } else{
                count.erase(id);
            }
            res[i] = f.empty() ? 0 : f.rbegin()->first;
          }
          return res;
    }
};