class Solution {
public:
    using ll = long long;
    ll mod = 1e9 + 7;

    struct E {
        int k, residue, mul;
        bool z;
    };

    ll modPow(ll a, ll e) {
        ll r = 1;
        while (e > 0) {
            if (e & 1) {
                r = (r * a) % mod;
            }
            a = (a * a) % mod;
            e >>= 1;
        }
        return r;
    }

    int xorAfterQueries(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        int q = queries.size();
        if (n == 0) {
            return 0;
        }
        int b = min(n, (int)sqrt(max(1, q)) + 1);
        vector<vector<E>> add(n), rem(n);
        vector<ll> lf(n, 1);
        for (auto& i : queries) {
            int l = i[0], r = i[1], k = i[2];
            ll v = i[3] % mod;
            if (v < 0) {
                v += mod;
            }
            if (k <= b) {
                int rs = l % k;
                int last = l + ((r - l) / k) * k;
                ll sp = (ll)last + k;
                if (v == 0) {
                    add[l].push_back({k, rs, 0, true});
                    if (sp < n) {
                        rem[sp].push_back({k, rs, 0, true});
                    }
                } else {
                    add[l].push_back({k, rs, (int)v, false});
                    int inv = (int)modPow(v, mod - 2);
                    if (sp < n) {
                        rem[sp].push_back({k, rs, inv, false});
                    }
                }
            } else {
                for (int id = l; id <= r; id += k) {
                    lf[id] = (lf[id] * v) % mod;
                }
            }
        }
        vector<vector<int>> prod(b + 1), z(b + 1);
        for (int k{1}; k <= b; k++) {
            prod[k].assign(k, 1);
            z[k].assign(k, 0);
        }
        ll res{0};
        for (int i{0}; i < n; i++) {
            for (auto& j : rem[i]) {
                if (j.z) {
                    z[j.k][j.residue]--;
                } else {
                    prod[j.k][j.residue] = (ll)prod[j.k][j.residue] * j.mul % mod;
                }
            }
            for (auto& j : add[i]) {
                if (j.z) {
                    z[j.k][j.residue]++;
                } else {
                    prod[j.k][j.residue] = (ll)prod[j.k][j.residue] * j.mul % mod;
                }
            }
            ll sf{1};
            for (int j{1}; j <= b; j++) {
                int r = i % j;
                if (z[j][r] > 0) {
                    sf = 0;
                    break;
                }
                sf = (sf * prod[j][r]) % mod;
            }
            ll cur = nums[i] % mod;
            if (cur < 0) {
                cur += mod;
            }
            cur = (cur * lf[i]) % mod;
            cur = (cur * sf) % mod;
            res ^= cur;
        }
        return (int)res;
    }
};