class Solution {
public:
    vector<int> longestObstacleCourseAtEachPosition(vector<int>& obstacles) {
        int n = obstacles.size();
        vector<int>res(n);
        vector<int>tail;
        for(int i {0}; i < n; i++){
            int h = obstacles[i];
            auto it = upper_bound(tail.begin(), tail.end(), h);
            if (it == tail.end()){
                tail.push_back(h);
                res[i] = tail.size();
            } else{
                *it = h;
                res[i] = it - tail.begin() + 1;
            }
        }
        return res;
    }
};