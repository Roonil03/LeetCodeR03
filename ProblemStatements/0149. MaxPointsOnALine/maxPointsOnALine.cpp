class Solution {
public:
    int maxPoints(vector<vector<int>>& points) {
        int n = points.size();
        if (n <= 2)
        {
            return n;
        }
        int maxCount = 2;
        for (int i = 0; i < n && maxCount < n; i++) {
            unordered_map<double, int> slopes;
            for (int j = i + 1; j < n; j++) {
                double slope;
                if (points[j][0] == points[i][0]) {
                    slope = DBL_MAX;
                } else {
                    slope = (double)(points[j][1] - points[i][1]) / (points[j][0] - points[i][0]);
                    if (slope == 0){
                        slope = 0;
                    }
                }
                slopes[slope]++;
                maxCount = max(maxCount, slopes[slope] + 1);
            }
        }
        return maxCount;
    }
};