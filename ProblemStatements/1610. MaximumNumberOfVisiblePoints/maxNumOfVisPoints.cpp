class Solution {
public:
    int visiblePoints(vector<vector<int>>& points, int angle, vector<int>& location) {
        int a {0}, n = points.size();
        vector<double>ag;
        for(auto& i: points){
            int dx = i[0] - location[0], dy = i[1] - location[1];
            if(dx == 0 && dy == 0){
                a++;
            } else{
                double at = atan2(dy, dx) * 180 / M_PI;
                if(at < 0){
                    at += 360;
                }
                ag.push_back(at);
            }
        }
        sort(ag.begin(), ag.end());
        vector<double> temp = ag;
        ag.insert(ag.end(), temp.begin(), temp.end());
        int res {0}, j {0};
        for(int i = n - a; i < ag.size(); i++){
            ag[i] += 360;
        }
        for(int i {0}; i < ag.size(); i++){
            while(j < ag.size() && ag[j] - ag[i] <= angle + 1e-6){
                j++;
            }
            res = max(res, j - i);
        }
        return min(res + a, n);
    }
};