class Solution {
public:
    using ll = long long;
    int threeSumMulti(vector<int>& arr, int target) {
        int mod = 1e9 + 7;
        sort(arr.begin(), arr.end());
        int n = arr.size();
        ll res {0};
        for(int i {0}; i < n; i++){
            int j = i + 1;
            int k = n - 1;
            ll a = (ll)target - arr[i];
            while(j < k){
                ll sum = (ll)arr[j] + (ll)arr[k];
                if(sum  < a){
                    j++;
                } else if(sum > a){
                    k--;
                } else{
                    if(arr[j] != arr[k]){
                        ll l {1}, r{1};
                        while(j + 1 < k && arr[j] == arr[j+1]){
                            l++;
                            j++;
                        }
                        while(k - 1 > j && arr[k] == arr[k-1]){
                            r++;
                            k--;
                        }
                        res = (res + l * r) % mod;
                        j++;
                        k--;
                    } else{
                        ll c = k - j + 1;
                        res = (res + c * (c - 1) / 2) % mod;
                        break;
                    }
                }
            }
        }
        return (int)(res % mod);
    }
};