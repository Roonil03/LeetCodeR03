class Solution {
public:
    vector<int> survivedRobotsHealths(vector<int>& positions, vector<int>& healths, string directions) {
        int n = positions.size();
        struct Robot {
            int pos;
            int health;
            char dir;
            int id;
        };
        vector<Robot> r(n);
        for (int i = 0; i < n; i++) {
            r[i] = {positions[i], healths[i], directions[i], i};
        }
        sort(r.begin(), r.end(), [](const Robot& a, const Robot& b) {
            return a.pos < b.pos;
        });
        vector<int> st;
        for (int i = 0; i < n; i++) {
            if (r[i].dir == 'R') {
                st.push_back(i);
                continue;
            }
            while (!st.empty() && r[i].health > 0) {
                int j = st.back();
                if (r[j].health < r[i].health) {
                    st.pop_back();
                    r[i].health--;
                    r[j].health = 0;
                } else if (r[j].health > r[i].health) {
                    r[j].health--;
                    r[i].health = 0;
                } else {
                    st.pop_back();
                    r[j].health = 0;
                    r[i].health = 0;
                }
            }
        }
        vector<pair<int,int>> surv;
        for (const auto& r : r) {
            if (r.health > 0) {
                surv.push_back({r.id, r.health});
            }
        }
        sort(surv.begin(), surv.end());
        vector<int> res;
        res.reserve(surv.size());
        for (auto& [id, hp] : surv) {
            res.push_back(hp);
        }
        return res;
    }
};
