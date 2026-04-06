class Solution {
public:
    using ll = long long;
    using pii = pair<int, int>;

    ll encode(int x, int y){
        return (static_cast<long long>(x) << 32) ^ static_cast<unsigned int>(y);
    }

    int robotSim(vector<int>& commands, vector<vector<int>>& obstacles) {
        pii pos;
        unordered_set<ll> mp;
        for(auto& i : obstacles){
            mp.insert(encode(i[0], i[1]));
        }
        int dx[4] = {0, 1, 0, -1};
        int dy[4] = {1, 0, -1, 0};
        int d {0};
        int res {0};
        for(int i : commands){
            if(i == -2){
                d = (d + 3) % 4;
            } else if(i == -1){
                d = (d + 1) % 4;
            } else{
                for(int j {0}; j < i; j++){
                    int a = pos.first + dx[d];
                    int b = pos.second + dy[d];
                    if(mp.count(encode(a, b))){
                        break;
                    }
                    pos.first = a;
                    pos.second = b;
                    res = max(res, pos.first * pos.first + pos.second * pos.second);
                }
            }
        }
        return res;
    }
};