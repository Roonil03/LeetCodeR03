class Solution {
public:
    bool isRectangleCover(vector<vector<int>>& rectangles) {
        int minX = INT_MAX, minY = INT_MAX;
        int maxX = INT_MIN, maxY = INT_MIN;
        long long area {0};
        set<pair<int,int>> points;
        for(auto& rect : rectangles){
            int x1 = rect[0], y1 = rect[1], x2 = rect[2], y2 = rect[3];
            minX = min(minX, x1);
            minY = min(minY, y1);
            maxX = max(maxX, x2);
            maxY = max(maxY, y2);
            area += (long long)(x2 - x1) * (y2 - y1);
            pair<int,int> p1 = {x1, y1};
            pair<int,int> p2 = {x1, y2};
            pair<int,int> p3 = {x2, y1};
            pair<int,int> p4 = {x2, y2};
            for(auto& p : {p1, p2, p3, p4}){
                if(points.count(p)) {
                    points.erase(p);
                } else {
                    points.insert(p);
                }
            }
        }
        if(points.size() != 4){
            return false;
        }
        if(!points.count({minX, minY})){
            return false;
        }
        if(!points.count({minX, maxY})){
            return false;
        }
        if(!points.count({maxX, minY})){
            return false;
        }
        if(!points.count({maxX, maxY})){
            return false;
        }
        long long temp = (long long)(maxX - minX) * (maxY - minY);
        return area == temp;
    }
};