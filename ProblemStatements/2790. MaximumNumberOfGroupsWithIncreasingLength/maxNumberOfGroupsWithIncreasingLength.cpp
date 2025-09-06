class Solution {
public:
    int maxIncreasingGroups(vector<int>& usageLimits) {
        // long long tot = accumulate(usageLimits.begin(), usageLimits.end(), 0LL);
        // int l {0}, r = usageLimits.size(), m;
        // auto check = [&](int a){
        //     long long sum = a * (a + 1) / 2;
        //     return sum <= tot;
        // };
        // while(l < r){
        //     m = l + (r - l + 1) / 2;
        //     if(check(m)){
        //         l = m;
        //     }
        //     else{
        //         r = m - 1;
        //     }
        // }
        // return l;

    //     int l {0}, r = usageLimits.size(), m;        
    //     auto check = [&](int g) -> bool{
    //         long long a {0};
    //         for (int t : usageLimits) {
    //             a += min(t, g);
    //         }
    //         long long h = (long long)g * (g + 1) / 2;
    //         return a >= h;
    //     };        
    //     while (l < r){
    //         m = l + (r - l + 1) / 2;
    //         if (check(m)){
    //             l = m;
    //         } else{
    //             r = m - 1;
    //         }
    //     }
    //     return l;
    // }

    int n = usageLimits.size();
        sort(usageLimits.begin(), usageLimits.end(), greater<int>());
        int l = 1, r = n;
        int res = 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (satisfy(mid, usageLimits)) {
                res = max(res, mid);
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }
        return res;
    }

    bool satisfy(int groups, vector<int> &usageLimits) {
        int gap = 0;
        int h = groups;
        for (int i = 0; i < usageLimits.size(); ++i) {
            gap = max(h - usageLimits[i] + gap, 0);
            if (h) h -= 1;
        }
        return gap == 0;
    }
};