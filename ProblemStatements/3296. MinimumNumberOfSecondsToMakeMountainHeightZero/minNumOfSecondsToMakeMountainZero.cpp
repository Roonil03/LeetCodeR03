class Solution {
public:
    using ll = long long;
    long long minNumberOfSeconds(int mountainHeight, vector<int>& workerTimes) {
        ll l {0};
        ll m = *min_element(workerTimes.begin(), workerTimes.end());
        ll r = m * (ll)mountainHeight * (mountainHeight + 1) / 2;
        ll res = r;
        while(l <= r){
            ll mid = l + (r - l) / 2;
            if(h1(mid, mountainHeight, workerTimes)){
                res = mid;
                r = mid - 1;
            } else{
                l = mid + 1;
            }
        }
        return res;
    }

    bool h1(ll t, int mountainHeight, vector<int>& workerTimes){
        ll tot {0};
        for(int i : workerTimes){
            ll lim = 2 * t / i;
            ll x = (sqrt(1 + 4 * lim) - 1) / 2;
            tot += x;
            if(tot >= mountainHeight){
                return true;
            }
        }
        return tot >= mountainHeight;
    }
};